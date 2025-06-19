package main

import (
	"context"
	"crypto/tls"
	"flag"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/emiago/sipgo"
	"github.com/emiago/sipgo/sip"
	sip_gb28181 "github.com/thinkgos/sip-gb28181"

	"net/http"
	_ "net/http/pprof"

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

	go func() {
		http.ListenAndServe(":6060", nil)
	}()

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

var chal digest.Challenge

func on_register(req *sip.Request, tx sip.ServerTransaction) {
	log := slog.Default()
	// https://www.rfc-editor.org/rfc/rfc2617#page-6
	h := req.GetHeader("Authorization")
	if h == nil {
		chal = digest.Challenge{
			Realm:     sip_domain,
			Nonce:     strconv.FormatInt(time.Now().UnixMicro(), 10),
			Algorithm: "MD5",
			QOP:       []string{"auth"},
		}

		res := sip.NewResponseFromRequest(req, 401, "Unauthorized", nil)
		res.AppendHeader(sip.NewHeader("WWW-Authenticate", chal.String()))
		tx.Respond(res)
		return
	}

	cred, err := digest.ParseCredentials(h.Value())
	if err != nil {
		log.Error("parsing creds failed", "error", err)
		tx.Respond(sip.NewResponseFromRequest(req, 401, "Bad credentials", nil))
		return
	}
	slog.Debug("parse cred", "cred", cred)
	// Check registry
	// passwd, exists := registry[cred.Username]
	// if !exists {
	// 	tx.Respond(sip.NewResponseFromRequest(req, 404, "Bad authorization header", nil))
	// 	return
	// }

	// Make digest and compare response
	digCred, err := digest.Digest(&chal, digest.Options{
		Method:   "REGISTER",
		URI:      cred.URI,
		Username: cred.Username,
		Password: sip_password,
		Count:    cred.Nc,
		A1:       "",
		Cnonce:   cred.Cnonce,
	})
	if err != nil {
		log.Error("Calc digest failed", "error", err)
		tx.Respond(sip.NewResponseFromRequest(req, 401, "Bad credentials", nil))
		return
	}
	if cred.Response != digCred.Response {
		log.Error("Calc digest failed", "req response", cred.Response, "cal response", digCred.Response)
		tx.Respond(sip.NewResponseFromRequest(req, 401, "Unauthorized", nil))
		return
	}
	log.Info("New client registered", "username", cred.Username)

	go func() {
		data, _ := sip_gb28181.MarshalXML(&sip_gb28181.CatalogQuery{
			CmdType:  "catalog",
			Sn:       1000,
			DeviceId: cred.Username,
		})
		req := sip.NewRequest(sip.MESSAGE, sip.Uri{
			Scheme:             "",
			Wildcard:           false,
			HierarhicalSlashes: false,
			User:               "",
			Password:           sip_password,
			Host:               "",
			Port:               0,
			UriParams:          sip.HeaderParams{},
			Headers:            sip.HeaderParams{},
		})
		req.SetBody(data)
		client.WriteRequest(req)
		if err != nil {
			log.Error("xxxxxxxxxx", "err", err)
		}
	}()
	tx.Respond(sip.NewResponseFromRequest(req, 200, "OK", nil))
}

func on_message(req *sip.Request, tx sip.ServerTransaction) {
	slog.Debug("on_message", "req", req)
	tx.Respond(sip.NewResponseFromRequest(req, 200, "ok", nil))
}
