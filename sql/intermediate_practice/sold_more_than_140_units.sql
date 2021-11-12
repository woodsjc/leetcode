SELECT DISTINCT
    s.name
FROM salesperson s
JOIN orders o
    on s.id = o.salesperson_id
GROUP BY s.name
HAVING sum(o.amount) > 1400
;
