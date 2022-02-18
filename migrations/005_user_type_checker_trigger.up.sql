CREATE OR REPLACE FUNCTION user_type_checker() RETURNS TRIGGER LANGUAGE PLPGSQL AS
$$
    BEGIN
        IF NEW.is_identified = true THEN
            INSERT INTO identifed_user_accounts(user_id) VALUES (NEW.id);
        ELSE
            INSERT INTO unidentifed_user_accounts(user_id) VALUES (NEW.id);
        END IF;
        RETURN NEW;
    END;
$$;

CREATE TRIGGER user_type_checker_trigger AFTER INSERT ON users
FOR EACH ROW EXECUTE PROCEDURE user_type_checker();

