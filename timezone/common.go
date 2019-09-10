package timezone

// NOTES:
//
// .Zone() returns both timezone and offset depending on the function, I ignore one of the two to keep the return simple.
//
// time.Now() defaults to UTC time, so the offset is going to always be 0 unless the timezone is specified
