// since this is standalone program (oposed to a library)
package main

type album struct {
	//this `json:"id" is for serialization`
	ID string `json:"id"`
}
