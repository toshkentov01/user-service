CREATE OR REPLACE FUNCTION password_hasher() RETURNS TRIGGER LANGUAGE PLPGSQL AS
$$
    BEGIN
        UPDATE users SET password = crypt(password, gen_salt('bf')),
            updated_at = NOW()
        WHERE id = NEW.id;

        RETURN NEW;
    END;
$$;

CREATE TRIGGER password_hasher_trigger AFTER INSERT ON users
FOR EACH ROW EXECUTE PROCEDURE password_hasher();