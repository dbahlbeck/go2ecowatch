![image >](assets/logo.png)

# ha2ecowatch

This is a simple Go app that adapts the EcoWatch MQTT API to a simpler MQTT interface more suitable for automation
products.

## Why?

There is an inner ring on the display of my EcoWatch Julia is unused, and I want to use the ECO watch to show more data
than just the hourly electricity price. So by adapting the message on the topic to the structure required by EcoWatch I
can easily display any arbitrary data from Home Assistant. In my case this is percentage of time remaining during 
a Bosch dishwasher program.

## Configuration

### Environment variables

- `G2E_HOST` MQTT server
- `G2E_PORT` MQTT port (default 1883)
- `G2E_USER` MQTT username
- `G2E_PASSWORD` MQTT password
- `G2E_ECOWATCH_ID` The ID of your EcoWatch, found on the web interface for your EcoWatch device.

### Config file

You could also configure go2ecowatch with a YAML config file in: `$HOME/.config/go2ecowatch.yaml`:

```yaml
user: <MQTT username>
password: <MQTT password>
host: <MQTT host>
port: <MQTT port (default 1883)>
ecowatch_id: <EcoWatch ID>
```

## Topics

### /go2ecowatch/inner/progressbar

Post a plain integer value to this topic to set a red to green gradient progress ring on the inner ring of the clock. The
adapted value (a JSON object) will be posted to `ecowatch/<ID>/set/pixels`

Invalid values will make the ring red.