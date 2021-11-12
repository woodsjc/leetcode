SELECT DISTINCT
    s.name
FROM customer c
JOIN orders o
    ON c.id = o.cust_id
JOIN salesperson s
    ON o.salesperson_id = s.id
WHERE c.name = 'Samsonic'
;
