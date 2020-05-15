// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/

package waitfor

import (
	"log"
	"time"

	"github.com/go-redis/redis/v7"
)

// Create a Redis client with the given options and then wait forever until it is alive. Keeps trying and never returns an error.
func Redis(options *redis.Options) (*redis.Client, error) {
        client := redis.NewClient(options)
        for {
                if _, err := client.Ping().Result(); err != nil {
			log.Printf("[W] Could to ping Redis at <redis://%s/%d>: %s", options.Addr, options.DB, err)
                        time.Sleep(5 * time.Second)
                        continue
                }
                return client, nil
        }
        return nil, nil
}
