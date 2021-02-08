CREATE OR REPLACE FUNCTION record_changed() RETURNS TRIGGER
    LANGUAGE plpgsql
    AS $$
BEGIN
  NEW.ModifiedDate := current_timestamp;
  
  RETURN NEW;
END;
$$;
