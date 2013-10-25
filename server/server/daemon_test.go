package main

import (
    "testing"
    "eos/server/test"
)

func Test_AuthoriseD_newOrg(t *testing.T) {
    setupOrg()
    org, err := GetOrg("Anonymous")
    test.Assert(err != nil, "organisation does not exist", t)

    d := &Daemon{Id: "test"}

    d.Authorise()
    org, err = GetOrg("Anonymous")

    test.Assert(org != nil, "It creates a new org", t)
    test.Assert(len(org.Daemons) == 1, "It adds daemon to organisation", t)
    test.Assert(d.OrgId == "Anonymous", "It stores org ID in daemon", t)
}

func Test_AuthoriseD_exOrg(t *testing.T) {
    setupOrg()
    org := NewOrg("Anonymous")
    test.Assert(len(org.Daemons) == 0, "there are no daemons in the org initially", t)

    d := &Daemon{Id: "test"}

    d.Authorise()

    test.Assert(len(org.Daemons) == 1, "It adds daemon to organisation", t)
    test.Assert(d.OrgId == "Anonymous", "It stores org ID in daemon", t)
}

func Test_DeauthoriseD(t *testing.T) {
    org := setupOrg()
    org.Daemons["test"] = &Connection{}

    d := &Daemon{Id: "test", OrgId: NO_ORG}

    err := d.Deauthorise()
    test.Assert(err != nil, "it does not deauthorise already deauthorised daemon", t)

    d.OrgId = "123"
    err = d.Deauthorise()
    test.Assert(err == nil && len(org.Daemons) == 0, "it removes a user from organisation", t)
    test.Assert(d.OrgId == NO_ORG, "it sets the user organisation to none", t)
}

func Test_IsAuthorisedD(t *testing.T) {
    d := &Daemon{OrgId:NO_ORG}
    test.Assert(!d.IsAuthorised(), "it spots unauthorised person", t)
    d.OrgId = "random"
    test.Assert(d.IsAuthorised(), "it tells when a person is authorised", t)
}