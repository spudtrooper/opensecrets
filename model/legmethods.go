package model

import (
	"bytes"
	"fmt"
)

// CID returns the string CID of this Legislator
func (l *Legislator) CID() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.CID, nil
}

// CID returns CID() or panics if there is an error
func (l *Legislator) CIDOrPanic() string {
	res, err := l.CID()
	if err != nil {
		panic(fmt.Sprintf("error calling CID(): %v", err))
	}
	return res
}

// Firstlast returns the string Firstlast of this Legislator
func (l *Legislator) Firstlast() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.Firstlast, nil
}

// Firstlast returns Firstlast() or panics if there is an error
func (l *Legislator) FirstlastOrPanic() string {
	res, err := l.Firstlast()
	if err != nil {
		panic(fmt.Sprintf("error calling Firstlast(): %v", err))
	}
	return res
}

// Lastname returns the string Lastname of this Legislator
func (l *Legislator) Lastname() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.Lastname, nil
}

// Lastname returns Lastname() or panics if there is an error
func (l *Legislator) LastnameOrPanic() string {
	res, err := l.Lastname()
	if err != nil {
		panic(fmt.Sprintf("error calling Lastname(): %v", err))
	}
	return res
}

// Party returns the string Party of this Legislator
func (l *Legislator) Party() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.Party, nil
}

// Party returns Party() or panics if there is an error
func (l *Legislator) PartyOrPanic() string {
	res, err := l.Party()
	if err != nil {
		panic(fmt.Sprintf("error calling Party(): %v", err))
	}
	return res
}

// Office returns the string Office of this Legislator
func (l *Legislator) Office() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.Office, nil
}

// Office returns Office() or panics if there is an error
func (l *Legislator) OfficeOrPanic() string {
	res, err := l.Office()
	if err != nil {
		panic(fmt.Sprintf("error calling Office(): %v", err))
	}
	return res
}

// Gender returns the string Gender of this Legislator
func (l *Legislator) Gender() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.Gender, nil
}

// Gender returns Gender() or panics if there is an error
func (l *Legislator) GenderOrPanic() string {
	res, err := l.Gender()
	if err != nil {
		panic(fmt.Sprintf("error calling Gender(): %v", err))
	}
	return res
}

// FirstElected returns the string FirstElected of this Legislator
func (l *Legislator) FirstElected() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.FirstElected, nil
}

// FirstElected returns FirstElected() or panics if there is an error
func (l *Legislator) FirstElectedOrPanic() string {
	res, err := l.FirstElected()
	if err != nil {
		panic(fmt.Sprintf("error calling FirstElected(): %v", err))
	}
	return res
}

// ExitCode returns the string ExitCode of this Legislator
func (l *Legislator) ExitCode() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.ExitCode, nil
}

// ExitCode returns ExitCode() or panics if there is an error
func (l *Legislator) ExitCodeOrPanic() string {
	res, err := l.ExitCode()
	if err != nil {
		panic(fmt.Sprintf("error calling ExitCode(): %v", err))
	}
	return res
}

// Comments returns the string Comments of this Legislator
func (l *Legislator) Comments() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.Comments, nil
}

// Comments returns Comments() or panics if there is an error
func (l *Legislator) CommentsOrPanic() string {
	res, err := l.Comments()
	if err != nil {
		panic(fmt.Sprintf("error calling Comments(): %v", err))
	}
	return res
}

// Phone returns the string Phone of this Legislator
func (l *Legislator) Phone() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.Phone, nil
}

// Phone returns Phone() or panics if there is an error
func (l *Legislator) PhoneOrPanic() string {
	res, err := l.Phone()
	if err != nil {
		panic(fmt.Sprintf("error calling Phone(): %v", err))
	}
	return res
}

// Fax returns the string Fax of this Legislator
func (l *Legislator) Fax() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.Fax, nil
}

// Fax returns Fax() or panics if there is an error
func (l *Legislator) FaxOrPanic() string {
	res, err := l.Fax()
	if err != nil {
		panic(fmt.Sprintf("error calling Fax(): %v", err))
	}
	return res
}

// Website returns the string Website of this Legislator
func (l *Legislator) Website() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.Website, nil
}

// Website returns Website() or panics if there is an error
func (l *Legislator) WebsiteOrPanic() string {
	res, err := l.Website()
	if err != nil {
		panic(fmt.Sprintf("error calling Website(): %v", err))
	}
	return res
}

// Webform returns the string Webform of this Legislator
func (l *Legislator) Webform() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.Webform, nil
}

// Webform returns Webform() or panics if there is an error
func (l *Legislator) WebformOrPanic() string {
	res, err := l.Webform()
	if err != nil {
		panic(fmt.Sprintf("error calling Webform(): %v", err))
	}
	return res
}

// CongressOffice returns the string CongressOffice of this Legislator
func (l *Legislator) CongressOffice() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.CongressOffice, nil
}

// CongressOffice returns CongressOffice() or panics if there is an error
func (l *Legislator) CongressOfficeOrPanic() string {
	res, err := l.CongressOffice()
	if err != nil {
		panic(fmt.Sprintf("error calling CongressOffice(): %v", err))
	}
	return res
}

// BioguideID returns the string BioguideID of this Legislator
func (l *Legislator) BioguideID() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.BioguideID, nil
}

// BioguideID returns BioguideID() or panics if there is an error
func (l *Legislator) BioguideIDOrPanic() string {
	res, err := l.BioguideID()
	if err != nil {
		panic(fmt.Sprintf("error calling BioguideID(): %v", err))
	}
	return res
}

// VotesmartID returns the string VotesmartID of this Legislator
func (l *Legislator) VotesmartID() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.VotesmartID, nil
}

// VotesmartID returns VotesmartID() or panics if there is an error
func (l *Legislator) VotesmartIDOrPanic() string {
	res, err := l.VotesmartID()
	if err != nil {
		panic(fmt.Sprintf("error calling VotesmartID(): %v", err))
	}
	return res
}

// FeccandID returns the string FeccandID of this Legislator
func (l *Legislator) FeccandID() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.FeccandID, nil
}

// FeccandID returns FeccandID() or panics if there is an error
func (l *Legislator) FeccandIDOrPanic() string {
	res, err := l.FeccandID()
	if err != nil {
		panic(fmt.Sprintf("error calling FeccandID(): %v", err))
	}
	return res
}

// TwitterID returns the string TwitterID of this Legislator
func (l *Legislator) TwitterID() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.TwitterID, nil
}

// TwitterID returns TwitterID() or panics if there is an error
func (l *Legislator) TwitterIDOrPanic() string {
	res, err := l.TwitterID()
	if err != nil {
		panic(fmt.Sprintf("error calling TwitterID(): %v", err))
	}
	return res
}

// YoutubeURL returns the string YoutubeURL of this Legislator
func (l *Legislator) YoutubeURL() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.YoutubeURL, nil
}

// YoutubeURL returns YoutubeURL() or panics if there is an error
func (l *Legislator) YoutubeURLOrPanic() string {
	res, err := l.YoutubeURL()
	if err != nil {
		panic(fmt.Sprintf("error calling YoutubeURL(): %v", err))
	}
	return res
}

// FacebookID returns the string FacebookID of this Legislator
func (l *Legislator) FacebookID() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.FacebookID, nil
}

// FacebookID returns FacebookID() or panics if there is an error
func (l *Legislator) FacebookIDOrPanic() string {
	res, err := l.FacebookID()
	if err != nil {
		panic(fmt.Sprintf("error calling FacebookID(): %v", err))
	}
	return res
}

// Birthdate returns the string Birthdate of this Legislator
func (l *Legislator) Birthdate() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.Birthdate, nil
}

// Birthdate returns Birthdate() or panics if there is an error
func (l *Legislator) BirthdateOrPanic() string {
	res, err := l.Birthdate()
	if err != nil {
		panic(fmt.Sprintf("error calling Birthdate(): %v", err))
	}
	return res
}

func (l *Legislator) String() string {
	var buf bytes.Buffer
	buf.WriteString("Legislator(")
	buf.WriteString("CID=")
	buf.WriteString(l.CIDOrPanic())
	buf.WriteString(", Firstlast=")
	buf.WriteString(l.FirstlastOrPanic())
	buf.WriteString(", Lastname=")
	buf.WriteString(l.LastnameOrPanic())
	buf.WriteString(", Party=")
	buf.WriteString(l.PartyOrPanic())
	buf.WriteString(", Office=")
	buf.WriteString(l.OfficeOrPanic())
	buf.WriteString(", Gender=")
	buf.WriteString(l.GenderOrPanic())
	buf.WriteString(", FirstElected=")
	buf.WriteString(l.FirstElectedOrPanic())
	buf.WriteString(", ExitCode=")
	buf.WriteString(l.ExitCodeOrPanic())
	buf.WriteString(", Comments=")
	buf.WriteString(l.CommentsOrPanic())
	buf.WriteString(", Phone=")
	buf.WriteString(l.PhoneOrPanic())
	buf.WriteString(", Fax=")
	buf.WriteString(l.FaxOrPanic())
	buf.WriteString(", Website=")
	buf.WriteString(l.WebsiteOrPanic())
	buf.WriteString(", Webform=")
	buf.WriteString(l.WebformOrPanic())
	buf.WriteString(", CongressOffice=")
	buf.WriteString(l.CongressOfficeOrPanic())
	buf.WriteString(", BioguideID=")
	buf.WriteString(l.BioguideIDOrPanic())
	buf.WriteString(", VotesmartID=")
	buf.WriteString(l.VotesmartIDOrPanic())
	buf.WriteString(", FeccandID=")
	buf.WriteString(l.FeccandIDOrPanic())
	buf.WriteString(", TwitterID=")
	buf.WriteString(l.TwitterIDOrPanic())
	buf.WriteString(", YoutubeURL=")
	buf.WriteString(l.YoutubeURLOrPanic())
	buf.WriteString(", FacebookID=")
	buf.WriteString(l.FacebookIDOrPanic())
	buf.WriteString(", Birthdate=")
	buf.WriteString(l.BirthdateOrPanic())
	buf.WriteString(")")
	return buf.String()
}
