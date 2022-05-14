package model

// CID returns the string CID of this Legislator
func (l *Legislator) CID() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.CID, nil
}

// Firstlast returns the string Firstlast of this Legislator
func (l *Legislator) Firstlast() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.Firstlast, nil
}

// Lastname returns the string Lastname of this Legislator
func (l *Legislator) Lastname() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.Lastname, nil
}

// Party returns the string Party of this Legislator
func (l *Legislator) Party() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.Party, nil
}

// Office returns the string Office of this Legislator
func (l *Legislator) Office() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.Office, nil
}

// Gender returns the string Gender of this Legislator
func (l *Legislator) Gender() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.Gender, nil
}

// FirstElected returns the string FirstElected of this Legislator
func (l *Legislator) FirstElected() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.FirstElected, nil
}

// ExitCode returns the string ExitCode of this Legislator
func (l *Legislator) ExitCode() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.ExitCode, nil
}

// Comments returns the string Comments of this Legislator
func (l *Legislator) Comments() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.Comments, nil
}

// Phone returns the string Phone of this Legislator
func (l *Legislator) Phone() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.Phone, nil
}

// Fax returns the string Fax of this Legislator
func (l *Legislator) Fax() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.Fax, nil
}

// Website returns the string Website of this Legislator
func (l *Legislator) Website() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.Website, nil
}

// Webform returns the string Webform of this Legislator
func (l *Legislator) Webform() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.Webform, nil
}

// CongressOffice returns the string CongressOffice of this Legislator
func (l *Legislator) CongressOffice() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.CongressOffice, nil
}

// BioguideID returns the string BioguideID of this Legislator
func (l *Legislator) BioguideID() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.BioguideID, nil
}

// VotesmartID returns the string VotesmartID of this Legislator
func (l *Legislator) VotesmartID() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.VotesmartID, nil
}

// FeccandID returns the string FeccandID of this Legislator
func (l *Legislator) FeccandID() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.FeccandID, nil
}

// TwitterID returns the string TwitterID of this Legislator
func (l *Legislator) TwitterID() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.TwitterID, nil
}

// YoutubeURL returns the string YoutubeURL of this Legislator
func (l *Legislator) YoutubeURL() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.YoutubeURL, nil
}

// FacebookID returns the string FacebookID of this Legislator
func (l *Legislator) FacebookID() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.FacebookID, nil
}

// Birthdate returns the string Birthdate of this Legislator
func (l *Legislator) Birthdate() (string, error) {
	info, err := l.getInfo()
	if err != nil {
		return "", err
	}
	return info.Birthdate, nil
}
