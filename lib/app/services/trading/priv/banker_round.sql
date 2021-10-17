-- round-half-even aka Banker's rounding
CREATE OR REPLACE FUNCTION 
    round(val NUMERIC, prec INT)
    returns NUMERIC

LANGUAGE 'plpgsql'
AS $$
DECLARE
    retval NUMERIC;
    difference NUMERIC;
    even BOOLEAN;
BEGIN
    retval := round(val,prec);
    difference := retval-val;
    IF abs(difference)*(10::NUMERIC^prec) = 0.5::NUMERIC THEN
        even := (retval * (10::NUMERIC^prec)) % 2::NUMERIC = 0::NUMERIC;
        IF NOT even THEN
            retval := round(val-difference,prec);
        END IF;
    END IF;
    RETURN retval;
END;
$$;