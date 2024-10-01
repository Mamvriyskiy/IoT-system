-- +goose Up
-- +goose StatementBegin
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
    RETURN -2; 
  END IF;

  IF current_status = 'inactive' AND newStatus = 'inactive' THEN
    RETURN -3; 
  END IF;

  UPDATE device
  SET status = newStatus
  WHERE deviceid = device_id;

  RETURN 0;
END;
$$ LANGUAGE plpgsql;

CREATE FUNCTION getHomeID(id integer, level integer, nameHome varchar(30)) 
RETURNS integer AS $$
DECLARE
  homeid integer;
BEGIN
  select h.homeid into homeid from home h 
	where h.homeid in (select a.homeid from access a
		where a.clientid = id) and h.name = nameHome;
  
  RETURN homeid;
END;
$$ LANGUAGE plpgsql;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION update_status(device_id integer, newStatus varchar(10));
DROP FUNCTION getHomeID(id integer, level integer, nameHome varchar(30));
-- +goose StatementEnd
