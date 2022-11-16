import http from "k6/http";
import { check, group, sleep } from "k6";
import { Counter, Rate, Trend } from "k6/metrics";
import { randomIntBetween } from "https://jslib.k6.io/k6-utils/1.0.0/index.js";

const loginData = JSON.parse(open("./users.json"));  // download the data file here: https://test.k6.io/static/examples/users.json

/* Options
Global options for your script
stages - Ramping pattern
thresholds - pass/fail criteria for the test
ext - Options used by Load Impact cloud service test name and distribution
*/
export let options = {

};


/* Main function
The main function is what the virtual users will loop over during test execution.
*/
export default function() {
    // We define our first group. Pages naturally fit a concept of a group
    // You may have other similar actions you wish to "group" together

        let res = null;
        // As mentioned above, this logic just forces the performance alert for too many urls, use env URL_ALERT to force it
        // It also highlights the ability to programmatically do things right in your script
        if (__ENV.URL_ALERT) {
            res = http.get("http://test.k6.io/?ts=" + Math.round(randomIntBetween(1,2000)));
        } else {
            res = http.get("http://test.k6.io/?ts=" + Math.round(randomIntBetween(1,2000)), { tags: { name: "http://test.k6.io/ Aggregated"}});
        }
        console.log(randomIntBetween(1,2000))
        // let checkRes = check(res, {
        //     "Homepage body size is 11026 bytes": (r) => r.body.length === 11026,
        //     "Homepage welcome header present": (r) => r.body.indexOf("Welcome to the k6.io demo site!") !== -1
        // });


}