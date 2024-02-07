package model

import "errors"


// ContactFormSubmission represents the structure of the JSON object
type ContactFormSubmission struct {
    Ev      string `json:"ev"`   // Event - Indicates the type of event
    Et      string `json:"et"`   // Event Type - Specifies the type of event
    Id      string `json:"id"`   // App ID - Identifier for the application
    Uid     string `json:"uid"`  // User ID - Unique identifier for the user
    Mid     string `json:"mid"`  // Message ID - Identifier for the submitted message
    T       string `json:"t"`    // Page Title - Title of the webpage where the form was submitted
    P       string `json:"p"`    // Page URL - URL of the webpage where the form was submitted
    L       string `json:"l"`    // Browser Language - Language of the browser used by the user
    Sc      string `json:"sc"`   // Screen Size - Resolution of the user's screen
    Atrk1   string `json:"atrk1"` // Attribute Key 1 - Key of the first attribute
    Atrv1   string `json:"atrv1"` // Attribute Value 1 - Value of the first attribute
    Atrt1   string `json:"atrt1"` // Attribute Type 1 - Type of the first attribute
    Atrk2   string `json:"atrk2"` // Attribute Key 2 - Key of the second attribute
    Atrv2   string `json:"atrv2"` // Attribute Value 2 - Value of the second attribute
    Atrt2   string `json:"atrt2"` // Attribute Type 2 - Type of the second attribute
    Uatrk1  string `json:"uatrk1"` // User Trait Key 1 - Key of the first user trait
    Uatrv1  string `json:"uatrv1"` // User Trait Value 1 - Value of the first user trait
    Uatrt1  string `json:"uatrt1"` // User Trait Type 1 - Type of the first user trait
    Uatrk2  string `json:"uatrk2"` // User Trait Key 2 - Key of the second user trait
    Uatrv2  string `json:"uatrv2"` // User Trait Value 2 - Value of the second user trait
    Uatrt2  string `json:"uatrt2"` // User Trait Type 2 - Type of the second user trait
    Uatrk3  string `json:"uatrk3"` // User Trait Key 3 - Key of the third user trait
    Uatrv3  string `json:"uatrv3"` // User Trait Value 3 - Value of the third user trait
    Uatrt3  string `json:"uatrt3"` // User Trait Type 3 - Type of the third user trait
}

// Validate function validates all fields of ContactFormSubmission struct
func (c *ContactFormSubmission) Validate() error {
    if c.Ev == "" {
        return errors.New("Event is required")
    }

    if c.Et == "" {
        return errors.New("Event Type is required")
    }

    if c.Id == "" {
        return errors.New("App ID is required")
    }

    if c.Uid == "" {
        return errors.New("User ID is required")
    }

    if c.Mid == "" {
        return errors.New("Message ID is required")
    }

    if c.T == "" {
        return errors.New("Page Title is required")
    }

    if c.P == "" {
        return errors.New("Page URL is required")
    }

    if c.L == "" {
        return errors.New("Browser Language is required")
    }

    if c.Sc == "" {
        return errors.New("Screen Size is required")
    }

    if c.Atrk1 == "" {
        return errors.New("Attribute Key 1 is required")
    }

    if c.Atrv1 == "" {
        return errors.New("Attribute Value 1 is required")
    }

    if c.Atrt1 == "" {
        return errors.New("Attribute Type 1 is required")
    }

    if c.Atrk2 == "" {
        return errors.New("Attribute Key 2 is required")
    }

    if c.Atrv2 == "" {
        return errors.New("Attribute Value 2 is required")
    }

    if c.Atrt2 == "" {
        return errors.New("Attribute Type 2 is required")
    }

    if c.Uatrk1 == "" {
        return errors.New("User Trait Key 1 is required")
    }

    if c.Uatrv1 == "" {
        return errors.New("User Trait Value 1 is required")
    }

    if c.Uatrt1 == "" {
        return errors.New("User Trait Type 1 is required")
    }

    if c.Uatrk2 == "" {
        return errors.New("User Trait Key 2 is required")
    }

    if c.Uatrv2 == "" {
        return errors.New("User Trait Value 2 is required")
    }

    if c.Uatrt2 == "" {
        return errors.New("User Trait Type 2 is required")
    }

    if c.Uatrk3 == "" {
        return errors.New("User Trait Key 3 is required")
    }

    if c.Uatrv3 == "" {
        return errors.New("User Trait Value 3 is required")
    }

    if c.Uatrt3 == "" {
        return errors.New("User Trait Type 3 is required")
    }

    return nil
}
