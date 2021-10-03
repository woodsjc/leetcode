DECLARE @n int = 1;

WHILE(@n <= 20)
BEGIN
    SELECT trim(replicate('* ', @n));
    SET @n = @n + 1;
END;
