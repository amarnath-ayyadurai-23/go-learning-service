package main

import "testing"

func TestHello(t *testing.T){

    got := basic()
    want := "\nHello, World!"

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestHypo(t *testing.T){

    res := basic()
    assert := "Hello, World!"

    if res != assert {
        t.Errorf("Not Equal: %q <> %q",res,assert)
    }

}