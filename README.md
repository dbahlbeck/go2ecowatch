![image >](assets/logo.png)

# go2ecowatch

This is a simple Go app that adapts the EcoWatch MQTT API to a simpler MQTT interface more suitable for automation products.

Note: work in progress!

## Why?

There is an inner ring on the display of my EcoWatch Julia that is unused, and I want to use the ECO watch to show more data
than just the hourly electricity price. So by adapting the message on the topic to the structure required by EcoWatch I
can easily display any arbitrary data from Home Assistant. In my case this is percentage of time remaining during 
a Bosch dishwasher program.

## Running

### From the code

Clone the repo, configure the environment variables, and run with `go run .`

### Docker

There's an image on docker hub: `wursley/go2ecowatch`. Pull it and run it with the necessary environment variable configuration.

### Runtipi

If you are using [Runtipi](https://runtipi.io/), then you could just add my AppStore [https://github.com/dbahlbeck/my-store](https://github.com/dbahlbeck/my-store) and install it to your home server directly from there.


## Try it out

When you have the server running, try pushing an int (0 - 100) to the topic `/go2ecowatch/inner/progressbar`

## Configuration

For this to work, you must have configured your EcoWatch to connect to your MQTT server. Go to the source to learn how to do that: [Waltrix support](https://waltrix.se/sv/pages/support)

### Environment variables

- `G2E_HOST` MQTT server
- `G2E_PORT` MQTT port (default: 1883)
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

### go2ecowatch/inner/progressbar

Post a plain integer value to this topic to set a red to green gradient progress ring on the inner ring of the clock. The
adapted value (a JSON object) will be posted to `ecowatch/<ID>/set/pixels`

Invalid values will make the ring red.