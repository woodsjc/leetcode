WITH hacker_submissions AS (
	SELECT
		h.hacker_id
		, h.name
		, s.submission_date
		, count(*) as total
		, row_number() over (
			partition by h.hacker_id, s.submission_date
			order by s.submission_date ASC, count(*) desc, h.hacker_id asc
		) as rn
	FROM submissions s
	JOIN hackers h 
		ON s.hacker_id = h.hacker_id
	GROUP BY h.hacker_id, h.name, s.submission_date
	--ORDER BY s.submission_date ASC, 3 DESC, h.hacker_id ASC
)
SELECT
	s.submission_date
	, count(DISTINCT s.hacker_id) AS total_hackers
	
	, (
		SELECT hs.hacker_id
		FROM hacker_submissions hs
		WHERE s.submission_date = hs.submission_date
			AND hs.rn = 1
	) as hacker_id
	
	
	, (
		SELECT hs.name
		FROM hacker_submissions hs
		WHERE s.submission_date = hs.submission_date
			AND hs.rn = 1
	) as name
FROM submissions s
JOIN hackers h
	ON s.hacker_id = h.hacker_id
JOIN hacker_submissions hs
	ON s.submission_date = hs.submission_date
	AND rn = 1
GROUP BY s.submission_date
ORDER BY s.submission_date ASC
;