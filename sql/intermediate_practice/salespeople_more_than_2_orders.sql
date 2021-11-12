SELECT
    s.name
FROM salesperson s
JOIN orders o
    on s.id = o.salesperson_id
GROUP by s.name
HAVING count(*) > 1
;
