package cst

const (
	//	STRAIN_SQL = `SELECT
	//	sg.strain_id,
	//	sg.number,
	//	sg.strain_name,
	//	sg.strain_annotate,
	//	sg.extra_key strain_extra_key,
	//	sg.extra_value strain_extra_value,
	//	sg.short_name,
	//	GROUP_CONCAT( sg.allele_name, '||' ) allele_name,
	//	GROUP_CONCAT( sg.allele_annotate, '||' ) allele_annotate,
	//	GROUP_CONCAT( sg.allele_extra_key, '||' ) allele_extra_key,
	//	GROUP_CONCAT( sg.allele_extra_value, '||' ) allele_extra_value,
	//	GROUP_CONCAT( sg.genome_name, '||' ) genome_name,
	//	GROUP_CONCAT( c.serial, '||' ) serial
	//FROM
	//	(
	//	SELECT
	//		strain_allele.*,
	//		g.genome_name,
	//		g.id gid
	//	FROM
	//		(
	//		SELECT
	//			sanana.*,
	//			GROUP_CONCAT( ane.extra_key, '||' ) allele_extra_key,
	//			GROUP_CONCAT( ane.extra_value, '||' ) allele_extra_value
	//		FROM
	//			(
	//			SELECT
	//				san.*,
	//				GROUP_CONCAT( ana.annotate, '||' ) allele_annotate
	//			FROM
	//				(
	//				SELECT
	//					an.id allele_id,
	//					ssasesn.*,
	//					GROUP_CONCAT( an.allele_name, '||' ) allele_name
	//				FROM
	//					(
	//					SELECT
	//						ssase.*,
	//						GROUP_CONCAT( sn.short_name, '||' ) short_name
	//					FROM
	//						(
	//						SELECT
	//							ssa.*,
	//							GROUP_CONCAT( se.extra_key, '||' ) extra_key,
	//							GROUP_CONCAT( se.extra_value, '||' ) extra_value
	//						FROM
	//							(
	//							SELECT
	//								s.id strain_id,
	//								s.number,
	//								s.strain_name,
	//								GROUP_CONCAT( sa.annotate, '||' ) strain_annotate
	//							FROM
	//								strain s
	//								LEFT JOIN strain_annotate sa ON s.id = sa.strain_id
	//							GROUP BY
	//								s.id
	//							) ssa
	//							LEFT JOIN strain_extra se ON se.strain_id = ssa.strain_id
	//						GROUP BY
	//							ssa.strain_id
	//						) ssase
	//						LEFT JOIN short_name sn ON sn.strain_id = ssase.strain_id
	//					GROUP BY
	//						ssase.strain_id
	//					) ssasesn
	//					LEFT JOIN allele_name an ON an.strain_id = ssasesn.strain_id
	//				GROUP BY
	//					an.id
	//				) san
	//				LEFT JOIN allele_name_annotate ana ON ana.allele_name_id = san.allele_id
	//			GROUP BY
	//				san.allele_id
	//			) sanana
	//			LEFT JOIN allele_name_extra ane ON sanana.allele_id = ane.allele_name_id
	//		GROUP BY
	//			sanana.allele_id
	//		) strain_allele
	//		LEFT JOIN genome g ON g.allel_name_id = strain_allele.allele_id
	//	) sg
	//	LEFT JOIN chromosome c ON sg.gid = c.genome_id
	//`

	//
	//	STRAIN_ALLELE_SQL = `SELECT
	//	sg.strain_id
	//FROM
	//	(
	//	SELECT
	//		strain_allele.*,
	//		g.genome_name,
	//		g.id gid
	//	FROM
	//		(
	//		SELECT
	//			sanana.*,
	//			GROUP_CONCAT( ane.extra_key, '||' ) allele_extra_key,
	//			GROUP_CONCAT( ane.extra_value, '||' ) allele_extra_value
	//		FROM
	//			(
	//			SELECT
	//				san.*,
	//				GROUP_CONCAT( ana.annotate, '||' ) allele_annotate
	//			FROM
	//				(
	//				SELECT
	//					an.id allele_id,
	//					ssasesn.*,
	//					GROUP_CONCAT( an.allele_name, '||' ) allele_name
	//				FROM
	//					(
	//					SELECT
	//						ssase.*,
	//						GROUP_CONCAT( sn.short_name, '||' ) short_name
	//					FROM
	//						(
	//						SELECT
	//							ssa.*,
	//							GROUP_CONCAT( se.extra_key, '||' ) extra_key,
	//							GROUP_CONCAT( se.extra_value, '||' ) extra_value
	//						FROM
	//							(
	//							SELECT
	//								s.id strain_id,
	//								s.number,
	//								s.strain_name,
	//								GROUP_CONCAT( sa.annotate, '||' ) strain_annotate
	//							FROM
	//								strain s
	//								LEFT JOIN strain_annotate sa ON s.id = sa.strain_id
	//							GROUP BY
	//								s.id
	//							) ssa
	//							LEFT JOIN strain_extra se ON se.strain_id = ssa.strain_id
	//						GROUP BY
	//							ssa.strain_id
	//						) ssase
	//						LEFT JOIN short_name sn ON sn.strain_id = ssase.strain_id
	//					GROUP BY
	//						ssase.strain_id
	//					) ssasesn
	//					LEFT JOIN allele_name an ON an.strain_id = ssasesn.strain_id
	//				GROUP BY
	//					an.id
	//				) san
	//				LEFT JOIN allele_name_annotate ana ON ana.allele_name_id = san.allele_id
	//			GROUP BY
	//				san.allele_id
	//			) sanana
	//			LEFT JOIN allele_name_extra ane ON sanana.allele_id = ane.allele_name_id
	//		GROUP BY
	//			sanana.allele_id
	//		) strain_allele
	//		LEFT JOIN genome g ON g.allel_name_id = strain_allele.allele_id
	//	) sg
	//	LEFT JOIN chromosome c ON sg.gid = c.genome_id
	//`
	//
	//
	//	STRAIN_ALLELE_LIKE_SQL = `	where
	//	(number like ? or strain_name like ? or strain_annotate like ?
	//		or strain_extra_key like ? or strain_extra_value like ? or allele_name like ?
	//		or allele_extra_key like ? or allele_extra_value like ? or allele_annotate like ?
	//		or genome_name like ? or serial like ? or  short_name like  ?  ) `
	//
	//
	//	STRAIN_ALLELE_END_SQL = `GROUP BY
	//	sg.strain_id`
	//
	//
	//	STRAIN_ALLELE_COUNT_SQL = `select count(*) from (%s)`

	STRAIN_SQL = `SELECT
	sasaese.*,
	GROUP_CONCAT(sn.short_name, '△' ) short_name 
FROM
	(
	SELECT
		sasae.*,
		GROUP_CONCAT( se.extra_key, '△' ) strain_extra_key,
		GROUP_CONCAT( se.extra_value, '△' ) strain_extra_value 
	FROM
		(
		SELECT
			sae.*,
			GROUP_CONCAT(sa.annotate,'△') strain_annotate 
		FROM
			(
			SELECT
				s.id,
				s.number,
				s.strain_name,
				GROUP_CONCAT( aaeaa.allele_id, '△' ) allele_id,
				GROUP_CONCAT( aaeaa.name, '△' ) allele_name,
				GROUP_CONCAT( aaeaa.genome, '△' ) genome,
				GROUP_CONCAT( aaeaa.serial, '△' ) serial,
				GROUP_CONCAT( aaeaa.extra_key, '△' ) allele_extra_key,
				GROUP_CONCAT( aaeaa.extra_value, '△' ) allele_extra_value,
				GROUP_CONCAT( aaeaa.annotate, '△' ) a_annotate 
			FROM
				strain s
				LEFT JOIN (
				SELECT
					aae.*,
					GROUP_CONCAT(aa.allele_id||'☆'||aa.annotate,'△' ) annotate 
				FROM
					(
					SELECT
						a.id allele_id,
						a.strain_id,
						a.name,
						a.genome,
						a.serial,
						GROUP_CONCAT( ae.allele_id||'☆'||ae.extra_key,'△' ) extra_key,
						GROUP_CONCAT( ae.allele_id||'☆'||ae.extra_value ,'△') extra_value 
					FROM
						allele a
						LEFT JOIN allele_extra ae ON ae.allele_id = a.id 
					GROUP BY
						a.id 
					) aae
					LEFT JOIN allele_annotate aa ON aa.allele_id = aae.allele_id 
				GROUP BY
					aae.allele_id 
				) aaeaa ON s.id = aaeaa.strain_id 
			GROUP BY
				s.id 
			) sae
			LEFT JOIN strain_annotate sa ON sae.id = sa.strain_id 
		GROUP BY
			sae.id 
		) sasae
		LEFT JOIN strain_extra se ON sasae.id = se.strain_id 
	GROUP BY
		sasae.id 
	) sasaese
	LEFT JOIN short_name sn ON sasaese.id = sn.strain_id 
`

	STRAIN_LIKE_SQL = `where
		(sasaese.number like ? or strain_name like ? or allele_name like ?
			or genome like ? or serial like ? or allele_extra_key like ? 
			or allele_extra_value like ? or a_annotate like ? or strain_annotate like ?
		or strain_extra_key like ? or  strain_extra_value like  ? or short_name like ? )`

	STRAIN_END_SQL = `  GROUP BY sasaese.id`

	STRAIN_COUNT_SQL = `select count(*) from (%s)`
)
