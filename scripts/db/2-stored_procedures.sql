\c catpix

CREATE OR REPLACE FUNCTION record_changed() RETURNS TRIGGER
    LANGUAGE plpgsql
    AS $$
BEGIN
  NEW.ModifiedDate := current_timestamp;
  
  RETURN NEW;
END;
$$;

CREATE TRIGGER trigger_record_changed
  BEFORE UPDATE ON catpix_user
  FOR EACH ROW
  EXECUTE PROCEDURE record_changed();

CREATE TRIGGER trigger_record_changed
  BEFORE UPDATE ON picture
  FOR EACH ROW
  EXECUTE PROCEDURE record_changed();
