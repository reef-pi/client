package cluster

import (
	"testing"
)

func TestAdmin(t *testing.T) {
	c := signIn(t)
	if _, err := c.Capabilities(); err != nil {
		t.Error(err)
	}
	if _, err := c.Info(); err != nil {
		t.Error(err)
	}
	if err := c.Reload(); err != nil {
		t.Error(err)
	}
	if err := c.PowerOff(); err != nil {
		t.Error(err)
	}
	if err := c.Reboot(); err != nil {
		t.Error(err)
	}

	if s, err := c.Settings(); err != nil {
		t.Error(err)
	} else {
		s.Interface = "en0"
		if err := c.UpdateSettings(s); err != nil {
			t.Error(err)
		}
	}
	if err := c.UpdateCredentials(Credentials{"reef-pi", "reef-pi"}); err != nil {
		t.Error(err)
	}

	if tc, err := c.Telemetry(); err != nil {
		t.Error(err)
	} else {
		tc.AdafruitIO.User = "client-test"
		if err := c.UpdateTelemetry(tc); err != nil {
			t.Error(err)
		}
	}

	if err := c.SendTestMessage(); err != nil {
		t.Error(err)
	}

	if _, err := c.Errors(); err != nil {
		t.Error(err)
	}

	if err := c.ClearErrors(); err != nil {
		t.Error(err)
	}

	if _, err := c.HealthStats(); err != nil {
		t.Error(err)
	}

	if ds, err := c.Dashboard(); err != nil {
		t.Error(err)
	} else {
		if err := c.UpdateDashboard(ds); err != nil {
			t.Error(err)
		}
	}

	if err := c.DisplayOn(); err != nil {
		t.Error(err)
	}

	if err := c.DisplayOff(); err != nil {
		t.Error(err)
	}

	if _, err := c.Display(); err != nil {
		t.Error(err)
	}

	if err := c.SignOut(); err != nil {
		t.Error(err)
	}
}
