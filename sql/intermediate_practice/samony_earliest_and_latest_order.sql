SELECT
    min(o.order_date) as earliest
    , max(o.order_date) as latest
FROM customer c
JOIN orders o
    ON c.id = o.cust_id
WHERE c.name = 'Samony'
;
