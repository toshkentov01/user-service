package utils

import "database/sql"

// StringToNullString ...
func StringToNullString(s string) (ns sql.NullString) {
	if s != "" {
		ns.String = s
		ns.Valid = true
	}
	return ns
}

//Float64ToNullFloat64 ...
func Float64ToNullFloat64(f float64) (nf sql.NullFloat64) {
	if f != 0 {
		nf.Float64 = f
		nf.Valid = true
	}
	return nf
}

// EmptyStringToNull also handles the case when user explicity wants to delete photo 
func EmptyStringToNull(s string) (res sql.NullString) {
	if s == "" {
		res.Valid = false
		return res
	} else if s == "delete"{
		res.String = ""
		res.Valid = true
		return res
	}else{
		res.String = s
		res.Valid = true
		return res
	}
}