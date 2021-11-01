package stringutils

import (
	"strings"
)

//Operations on String that are nil safe.
//IsEmpty/IsBlank - checks if a String contains text
//Trim/Strip - removes leading and trailing whitespace
//Equals - compares two strings nil-safe
//startsWith - check if a String starts with a prefix nil-safe
//endsWith - check if a String ends with a suffix nil-safe
//IndexOf/LastIndexOf/Contains - nil-safe index-of checks
//IndexOfAny/LastIndexOfAny/IndexOfAnyBut/LastIndexOfAnyBut - index-of any of a set of Strings
//ContainsOnly/ContainsNone/ContainsAny - does String contains only/none/any of these characters
//Substring/Left/Right/Mid - nil-safe substring extractions
//SubstringBefore/SubstringAfter/SubstringBetween - substring extraction relative to other strings
//Split/Join - splits a String into an array of substrings and vice versa
//Remove/Delete - removes part of a String
//Replace/Overlay - Searches a String and replaces one String with another
//Chomp/Chop - removes the last part of a String
//LeftPad/RightPad/Center/Repeat - pads a String
//UpperCase/LowerCase/SwapCase/Capitalize/Uncapitalize - changes the case of a String
//CountMatches - counts the number of occurrences of one String in another
//IsAlpha/IsNumeric/IsWhitespace/IsAsciiPrintable - checks the characters in a String
//DefaultString - protects against a nil input String
//Reverse/ReverseDelimited - reverses a String
//Abbreviate - abbreviates a string using ellipsis
//Difference - compares Strings and reports on their differences
//LevensteinDistance - the number of changes needed to change one String into another
//The StringUtils class defines certain words related to String handling.
//empty - a zero-length string ("")
//space - the space character (' ', char 32)
//whitespace - the characters defined by Character.isWhitespace(char)
//trim - the characters <= 32 as in String.trim()
//StringUtils handles nil input Strings quietly. That is to say that a nil input will return nil. Where a boolean or int is being returned details vary by method.
//A side effect of the nil handling is that a nilPointerException should be considered a bug in StringUtils (except for deprecated methods).
//Methods in this class give sample code to explain their operation. The symbol * is used to indicate any input including nil.
//#ThreadSafe#
//Since:
//		1.0
//See Also:
//		String
//Author:
//		Apache Software Foundation, Apache Jakarta Turbine , Jon S. Stevens, Daniel L. Rall, Greg Coladonato, Ed Korthof, Rand McNeely, Fredrik Westermarck, Holger Krauth, Alexander Day Chaffee, Henning P. Schmiedehausen, Arun Mammen Thomas, Gary Gregory, Phil Steitz, Al Chou, Michael Davey, Reuben Sivan, Chris Hyzer, Scott Johnson
//Version:
//		$Id:  $

// EMPTY The empty String "".
const EMPTY = ""

// WHITESPACE The whitespace String " ".
const WHITESPACE = " "

// INDEX_NOT_FOUND Represents a failed index search.
const INDEX_NOT_FOUND = -1

// PAD_LIMIT The maximum size to which the padding constant(s) can expand.
const PAD_LIMIT = 8192

// IsEmpty Checks if a String is empty ("").
//       IsEmpty("")        = true
//       IsEmpty(" ")       = false
//       IsEmpty("bob")     = false
//       IsEmpty("  bob  ") = false
//
//NOTE: That functionality is available in isBlank().
//Params:
//str – the String to check
//Returns:
//true if the String is empty
func IsEmpty(str string) bool {
	return len(str) == 0
}

// IsNotEmpty Checks if a String is not empty ("").
//       IsNotEmpty("")        = false
//       IsNotEmpty(" ")       = true
//       IsNotEmpty("bob")     = true
//       IsNotEmpty("  bob  ") = true
//	
//Params:
//str – the String to check
//Returns:
//true if the String is not empty
func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

// IsBlank Checks if a String is whitespace.
//       IsBlank("")        = true
//       IsBlank(" ")       = true
//       IsBlank("bob")     = false
//       IsBlank("  bob  ") = false
//	
//Params:
//str – the String to check
//Returns:
//true if the String is empty or whitespace
func IsBlank(str string) bool {
	return len(strings.Trim(str, WHITESPACE)) == 0
}

// IsNotBlank Checks if a String is not empty ("").
//       IsNotBlank("")        = false
//       IsNotBlank(" ")       = false
//       IsNotBlank("bob")     = true
//       IsNotBlank("  bob  ") = true
//
//Params:
//str – the String to check
//Returns:
//true if the String is not empty and not whitespace
func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

// Clean Removes control string from both ends of this String, handling by returning an empty String ("").
//       StringUtils.clean("")            = ""
//       StringUtils.clean("abc")         = "abc"
//       StringUtils.clean("    abc    ") = "abc"
//       StringUtils.clean("     ")       = ""
//
//Params:
//str – the String to clean
//Returns:
//the trimmed text
//See Also:
//strings.TrimSpace()
func Clean(str string) string {
	return strings.TrimSpace(str)
}

// Trim Removes control string from both ends of this String.
//The String is trimmed using strings.TrimSpace(). Trim removes start and end.
//       Trim("")            = ""
//       Trim("     ")       = ""
//       Trim("abc")         = "abc"
//       Trim("    abc    ") = "abc"
//		 Trim(" ab c ")      = "ab c"
//       
//Params:
//str – the String to be trimmed
func Trim(str string) string {
	return strings.TrimSpace(str)
}

// Equals Compares two Strings, returning true if they are equal.
//The comparison is case sensitive.
//       Equals("abc", "abc") = true
//       Equals("abc", "ABC") = false
//
//Params:
//str1 – the first String
//str2 – the second String
//Returns:
//true if the Strings are equal, case sensitive
func Equals(str1, str2 string) bool {
	return str1 == str2
}
