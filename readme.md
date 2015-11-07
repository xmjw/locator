# Michael Wawra, Technical Test

This application allows you to use Reverse Geocoding on Google and OpenStreetMap.

It can be run with:

`go run locator.go --lat=<latitude> --long=<longitude>`

e.g.: 

`go run locator.go --lat=52.5487429714954 --long=-1.81602098644987`

The application uses go channels to use both services at the same time, while also using a simple timeout to prevent it waiting indefinitely if neither service responds. (Note that Google is __almost__ always faster to respond.)

I did not have enough time to add tests on this occasion (although some infrastructure is in place).

Bitly's SimpleJson is the only non-core dependency, saved with `GoDep`.

## Additional Work

- I did add an interface definition for the two reverse geocoders, but it was unused to I removed it.
- Tests!
- I would like to add more reverse geocoders.

