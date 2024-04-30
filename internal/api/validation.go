package api

import (
    "errors"
    "net/url"
)

func validateSaveRequest(r *Request) error {
    if r.URL == "" {
        return errors.New("request URL is empty")
    }

    if _, err := url.ParseRequestURI(r.URL); err != nil {
        return errors.New("request URL is invalid")
    }

    return nil
}
