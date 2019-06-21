### Why?
If you have 1 or 2 tunnels or most of your tunnels are local ephemeral tunnels you may not need this.  However if you find yourself manually managing 2 or more tunnels, High Availability tunnels, or writing complex scripts to manage tunnel state this is the tool for you!


## Custom Tunnel Management
Perhaps you are a Unix Administrator who uses Cronjobs to keep a growing set of tunnels alive?  Maybe you have a set of bash scripts to manage tunnels:

```
# crontab -e myuser

# live-tunnel.bash
* * * * * bash live-tunnel.bash
# mac-tunnel.bash
* * * * * bash mac-tunnel.bash
# staging-tunnel.bash
* * * * * bash staging-tunnel.bash
# mobile-tunnel.bash
* * * * * bash mobile-tunnel.bash

# restart after utc noon 
0 */12 * * * bash /home/maxdobeck/tunnels/daily-restart.bash live-tunnel
0 */12 * * * bash /home/maxdobeck/tunnels/daily-restart.bash mac-tunnel
0 */12 * * * bash /home/maxdobeck/tunnels/daily-restart.bash staging-tunnel
0 */12 * * * bash /home/maxdobeck/tunnels/daily-restart.bash mobile-tunnel
```

If its not bash scripts then maybe its powershell scripts or .bat files.  You may or may not actually use the tunnels but the uptime and logs have become an additional duty.

## Easier Managed Tunnels
Use one configuration file defined by you or as many config files as you want.  With the `sauced` application you can start tunnels just like you did but with a centralized configuration.  Or manually like `sauced start ~/saucelabs/config.txt`.

Replacing the crontab from above!

```
# crontab -e myuser

* * * * * /Users/maxdobeck/Applications/sauced start /Users/maxdobeck/workspace/sauce_connect/sauced_test/sauced_config.txt >> /Users/maxdobeck/logs/    cron.log 2>&1
```

Once the tunnels are started you can check on them with `sauced show` and  see the metadata for your tunnels.

```
"pid": 2839,
"scbinary": "/Users/maxdobeck/workspace/sauce_connect/sc-4.5.1-osx/bin/sc",
"args": "/Users/maxdobeck/workspace/sauce_connect/sc-4.5.1-osx/bin/sc -u user.name -k your-access-key -v --no-remove-colliding-tunnels -N -i main-tunnel-pool --se-port 0 --pidfile /tmp/sc_client-2.pid",
"launchtime": "2019-06-21T22:12:58.612153Z",
"log": "/var/folders/b6/tp7rt8ws25bc_d9wjvmw0fc80000gq/T/sc-main-tunnel-pool.log",
"metadata": {
    "Pool": "main-tunnel-pool",
    "Size": 3,
    "Owner": "max.dobeck"
}
```

And when you want to stop them you can run `sauced stop` manually, on a schedule, or just pass a SIGINT or KILL command to a tunnel of your choice!



