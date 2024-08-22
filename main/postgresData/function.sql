DROP FUNCTION update_status(device_id integer, newStatus varchar(10)) ;

CREATE FUNCTION update_status(device_id integer, newStatus varchar(10)) 
RETURNS integer AS $$
DECLARE
  current_status varchar(10);
BEGIN
  SELECT status INTO current_status FROM device WHERE deviceid = device_id;

  IF NOT FOUND THEN
    RETURN -1; 
  END IF;

  IF current_status = 'active' AND newStatus = 'active' THEN
    RETURN -1; 
  END IF;

  UPDATE device
  SET status = newStatus
  WHERE deviceid = device_id;

  RETURN 0;
END;
$$ LANGUAGE plpgsql;

-- select update_status(1, 'Inactive');
