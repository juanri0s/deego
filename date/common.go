package date

// NOTES:
//
// For functions that accept strings, if your string doesn't follow the year-month-day ISO8601 layout,
// replace the layout string and it should still work.
//
// For functions that accept go time, if you want to use a string, then you will first need to convert
// using time.Parse.
//
// If you prefer accepting time over strings, then the conversion in some methods is not needed and only
// the logic needs to be used.
