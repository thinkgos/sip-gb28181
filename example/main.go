package main

import (
	"context"
	"crypto/tls"
	"flag"
	"log/slog"
	"math/rand/v2"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/emiago/sipgo"
	"github.com/emiago/sipgo/sip"
	sip_gb28181 "github.com/thinkgos/sip-gb28181"

	"github.com/icholy/digest"
)

const sip_id = "3402000000200000001"
const sip_domain = "3402000000"
const sip_password = "123456"

var client *sipgo.Client

func main() {
	extIP := flag.String("ip", ":15060", "My exernal ip")
	creds := flag.String("u", "3402000000200000001:123456", "Coma seperated username:password list")
	tran := flag.String("t", "udp", "Transport")
	tlskey := flag.String("tlskey", "", "TLS key path")
	tlscrt := flag.String("tlscrt", "", "TLS crt path")
	flag.Parse()

	// Make SIP Debugging available
	sip.SIPDebug = true //os.Getenv("SIP_DEBUG") != ""

	log := getLogger()

	registry := make(map[string]string)
	for _, c := range strings.Split(*creds, ",") {
		arr := strings.Split(c, ":")
		registry[arr[0]] = arr[1]
	}

	ua, err := sipgo.NewUA(
		sipgo.WithUserAgent("SIPGO"),
	// sipgo.WithUserAgentIP(*extIP),
	)
	if err != nil {
		log.Error("Fail to setup user agent", "error", err)
		return
	}

	client, err = sipgo.NewClient(ua, sipgo.WithClientAddr(*extIP))
	if err != nil {
		log.Error("Fail to setup server client", "error", err)
		return
	}

	srv, err := sipgo.NewServer(ua)
	if err != nil {
		log.Error("Fail to setup server handle", "error", err)
		return
	}

	// NOTE: This server only supports 1 REGISTRATION/Chalenge
	// This needs to be rewritten in better way
	srv.OnRegister(on_register)
	srv.OnMessage(on_message)
	log.Info("Listening on", "addr", *extIP)

	ctx := context.TODO()
	switch *tran {
	case "tls", "wss":
		cert, err := tls.LoadX509KeyPair(*tlscrt, *tlskey)
		if err != nil {
			log.Error("Fail to load  x509 key and crt", "error", err)
			return
		}
		if err := srv.ListenAndServeTLS(ctx, *tran, *extIP, &tls.Config{Certificates: []tls.Certificate{cert}}); err != nil {
			log.Info("Listening stop", "error", err)
		}
		return
	}

	if err := srv.ListenAndServe(ctx, *tran, *extIP); err != nil {
		log.Error("Failed to listen", "error", err)
	}
}

func getLogger() *slog.Logger {
	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro
	// zlog := zerolog.New(zerolog.ConsoleWriter{
	// 	Out:        os.Stdout,
	// 	TimeFormat: time.StampMicro,
	// }).With().Timestamp().Logger().Level(zerolog.InfoLevel)

	// logger := slog.New(slogzerolog.Option{Level: lvl, Logger: &zlog}.NewZerologHandler())
	// h := slog.NewTextHandler(os.Stdin, &slog.HandlerOptions{Level: lvl})
	// slog.SetDefault(slog.New(h))
	// var lvl slog.Level
	// if err := lvl.UnmarshalText([]byte(os.Getenv("LOG_LEVEL"))); err != nil {
	// 	lvl = slog.LevelInfo
	// }
	lvl := slog.LevelDebug
	slog.SetLogLoggerLevel(lvl)
	return slog.Default()
}

var cache_nonce = atomic.Pointer[string]{}

func init() {
	s := ""
	cache_nonce.Store(&s)
}

func on_register(req *sip.Request, tx sip.ServerTransaction) {
	log := slog.Default()
	// https://www.rfc-editor.org/rfc/rfc2617#page-6
	hdrAuth := req.GetHeader("Authorization")
	if hdrAuth == nil {
		nonce := strconv.FormatInt(time.Now().UnixMicro(), 10)
		cache_nonce.Store(&nonce)
		challenge := &digest.Challenge{
			Realm:     sip_domain,
			Nonce:     nonce,
			Algorithm: "MD5",
			QOP:       []string{"auth"},
		}
		res := sip.NewResponseFromRequest(req, 401, "Unauthorized", nil)
		res.AppendHeader(sip.NewHeader("WWW-Authenticate", challenge.String()))
		tx.Respond(res)
		return
	}

	cred, err := digest.ParseCredentials(hdrAuth.Value())
	if err != nil {
		log.Error("parsing authorization failure", "error", err)
		tx.Respond(sip.NewResponseFromRequest(req, 401, "Bad credentials", nil))
		return
	}
	slog.Debug("parse cred", "cred", cred)

	// Make digest and compare response
	digCred, err := digest.Digest(
		&digest.Challenge{
			Realm:     sip_domain,
			Nonce:     *cache_nonce.Load(),
			Algorithm: "MD5",
			QOP:       []string{"auth"},
		},
		digest.Options{
			Method:   "REGISTER",
			URI:      cred.URI,
			Username: cred.Username,
			Password: sip_password,
			Count:    cred.Nc,
			A1:       "",
			Cnonce:   cred.Cnonce,
		},
	)
	if err != nil {
		log.Error("Calc digest failed", "error", err)
		tx.Respond(sip.NewResponseFromRequest(req, 401, "Bad credentials", nil))
		return
	}
	if cred.Response != digCred.Response {
		log.Error("Calc digest not equal", "req response", cred.Response, "calc response", digCred.Response)
		tx.Respond(sip.NewResponseFromRequest(req, 401, "Unauthorized", nil))
		return
	}
	log.Info("New client registered", "username", cred.Username)

	req.Destination()
	host := req.Via().Host
	port := req.Via().Port
	hdrFrom := req.From()
	hdrTo := req.To()
	remoteAddr := req.Source()
	dest := req.Destination()
	go func() {
		data, _ := sip_gb28181.MarshalXML(&sip_gb28181.DeviceInfoQuery{
			CmdType:  sip_gb28181.CmdType_DeviceInfo,
			Sn:       int64(rand.IntN(800000) + 100000),
			DeviceId: cred.Username,
		})
		req := sip.NewRequest(sip.MESSAGE,
			sip.Uri{
				Scheme:             "sip",
				Wildcard:           false,
				HierarhicalSlashes: false,
				User:               cred.Username,
				Password:           sip_password,
				Host:               host,
				Port:               port,
				UriParams:          sip.HeaderParams{},
				Headers:            sip.HeaderParams{},
			})
		from := hdrTo.AsFrom()
		req.AppendHeader(&from)
		to := hdrFrom.AsTo()
		req.AppendHeader(&to)
		req.AppendHeader(&sip.ViaHeader{
			ProtocolName:    "SIP",
			ProtocolVersion: "2.0",
			Transport:       "UDP",
			Host:            host,
			Port:            port,
			Params: sip.NewParams().
				Add("branch", sip.GenerateBranch()),
		})
		req.SetSource(dest)
		req.SetDestination(remoteAddr)
		req.SetBody(data)
		resp, err := client.Do(context.Background(), req)

		if err != nil {
			log.Error("xxxxxxxxxx", "err", err)
		} else {
			log.Debug("message resp", "resp", resp)
		}
	}()
	tx.Respond(sip.NewResponseFromRequest(req, 200, "OK", nil))
}

func on_message(req *sip.Request, tx sip.ServerTransaction) {
	slog.Debug("on_message", "req", req)
	tx.Respond(sip.NewResponseFromRequest(req, 200, "ok", nil))
}
