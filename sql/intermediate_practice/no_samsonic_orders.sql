SELECT DISTINCT
    s.name
FROM salesperson s
WHERE NOT EXISTS (
    SELECT 1
    FROM customer c
    JOIN orders o
        ON c.id = o.cust_id
        AND o.salesperson_id = s.id
    WHERE c.name = 'Samsonic'
)
;
