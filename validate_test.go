package sip_gb28181

import "testing"

func Test_IsValidSipId(t *testing.T) {
	testCases := []struct {
		name   string
		domain string
		want   bool
	}{
		{
			name:   "少于20位",
			domain: "0123456789012345678",
			want:   false,
		},
		{
			name:   "含有非数字",
			domain: "0123456789012345678a",
			want:   false,
		},
		{
			name:   "正确",
			domain: "01234567890123456789",
			want:   true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsValidSipId(tc.domain); got != tc.want {
				t.Errorf("IsValidDomain(\"%s\") = %v, want %v", tc.domain, got, tc.want)
			}
		})
	}
}

func Test_IsValidSipDomainId(t *testing.T) {
	testCases := []struct {
		name   string
		domain string
		want   bool
	}{
		{
			name:   "少于10位",
			domain: "012345678",
			want:   false,
		},
		{
			name:   "含有非数字",
			domain: "012345678a",
			want:   false,
		},
		{
			name:   "正确",
			domain: "0123456789",
			want:   true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsValidSipDomain(tc.domain); got != tc.want {
				t.Errorf("IsValidDomain(\"%s\") = %v, want %v", tc.domain, got, tc.want)
			}
		})
	}
}
