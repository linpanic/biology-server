package cst

const (
	ALLELE_LIST_SQL = `SELECT
	* 
FROM
	(
	SELECT
		aae.*,
		GROUP_CONCAT( aa.annotate, '△' ) annotate 
	FROM
		(
		SELECT
			a.id allele_id,
			a.strain_id,
			a.name,
			a.genome,
			a.serial,
			GROUP_CONCAT( ae.extra_key, '△' ) extra_key,
			GROUP_CONCAT( ae.extra_value, '△' ) extra_value 
		FROM
			allele a
			LEFT JOIN allele_extra ae ON ae.allele_id = a.id 
		GROUP BY
			a.id 
		) aae
		LEFT JOIN allele_annotate aa ON aa.allele_id = aae.allele_id 
	GROUP BY
	aae.allele_id 
	)  `

	ALLELE_WHERE_SQL = ` where (name like ? or genome like ? or serial like ?  or annotate like ? or extra_key like ? or extra_value like ?)`
)
