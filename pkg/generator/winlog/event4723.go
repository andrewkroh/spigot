package winlog

import (
	"math/rand"
	"strconv"
)

const event4723 = 4723

// randomize4723 generates a random event with
// ID 4723 (An attempt was made to change an account's password).
func randomize4723(g *Generator) Event {
	domain := RandomDomain()
	hostname := RandomComputerName("")
	computerName := hostname + "." + domain

	targetName := hostname + "$"
	subjectName := RandomUser()

	evt := RandomEvent(event4723, g.getTime())
	evt.Provider = Provider{
		Name: "Microsoft-Windows-Security-Auditing",
		GUID: "{54849625-5478-4994-A5BA-3E3B0328C30D}",
	}
	evt.Channel = "Security"
	evt.Computer = computerName
	evt.EventData = EventData{
		Data: []KeyValue{
			{Key: "TargetUserSid", Value: RandomUserSID(targetName)},
			{Key: "TargetUserName", Value: targetName},
			{Key: "TargetDomainName", Value: domain},
			{Key: "SubjectUserSid", Value: RandomUserSID(subjectName)},
			{Key: "SubjectUserName", Value: subjectName},
			{Key: "SubjectDomainName", Value: domain},
			{Key: "SubjectLogonId", Value: "0x" + strconv.FormatInt(int64(rand.Intn(65536)), 16)},
			{Key: "PrivilegeList", Value: "-"},
		},
	}

	return evt
}
