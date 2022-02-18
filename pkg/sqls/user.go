package sqls

const (
	// ListSimilarUsers ...
	ListSimilarUsers = `
		WITH my_preferences AS  
			(SELECT skin_types, eye_types, hair_types FROM user_body_types WHERE user_id = $1)

		SELECT id, profile_photo, full_name, skin_types AS his_skin_types, eye_types AS his_eye_types, hair_types AS his_hair_types
		FROM user_client JOIN user_body_types ON id=user_id
		WHERE id != $1
		ORDER BY text_arr_similarity(skin_types, (SELECT skin_types FROM my_preferences)) + text_arr_similarity(eye_types, (SELECT eye_types FROM my_preferences)) + text_arr_similarity(hair_types, (SELECT hair_types FROM my_preferences)) DESC
	`

	// GetMyProfile ...
	GetMyProfile = `
		SELECT
			uc.username,
			ua.email,
			uc.full_name,
			uc.phone_number,
			uc.bio,
			uc.gender,
			ub.skin_types,
			ub.eye_types,
			ub.hair_types,
			ua.reg_type,
			uc.profile_photo
		FROM user_client AS uc 
		JOIN user_auth AS ua ON ua.id = uc.id
		JOIN user_body_types AS ub ON ub.user_id = uc.id
		WHERE ua.deleted_at IS NULL AND uc.id = $1
	`
)