package vultr

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	liveTest bool
	apiKey   string
	domain   string
	ip       string
)

func init() {
	apiKey = os.Getenv("VULTR_API_KEY")
	domain = os.Getenv("VULTR_TEST_DOMAIN")
	ip = os.Getenv("VULTR_TEST_IP")
	liveTest = len(apiKey) > 0 && len(domain) > 0
}

func restoreEnv() {
	os.Setenv("VULTR_API_KEY", apiKey)
}

func TestNewDNSProviderValidEnv(t *testing.T) {
	os.Setenv("VULTR_API_KEY", "123")
	defer restoreEnv()
	_, err := NewDNSProvider()
	assert.NoError(t, err)
}

func TestNewDNSProviderMissingCredErr(t *testing.T) {
	os.Setenv("VULTR_API_KEY", "")
	defer restoreEnv()
	_, err := NewDNSProvider()
	assert.EqualError(t, err, "Vultr credentials missing")
}

func TestLiveEnsureARecord(t *testing.T) {
	if !liveTest {
		t.Skip("skipping live test")
	}

	provider, err := NewDNSProvider()
	assert.NoError(t, err)

	err = provider.EnsureARecord(domain, ip)
	assert.NoError(t, err)
}

func TestLiveDeleteARecords(t *testing.T) {
	if !liveTest {
		t.Skip("skipping live test")
	}

	time.Sleep(time.Second * 1)

	provider, err := NewDNSProvider()
	assert.NoError(t, err)

	err = provider.DeleteARecords(domain)
	assert.NoError(t, err)
}
