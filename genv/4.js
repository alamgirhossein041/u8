
			var parseHTML = require("k6/html").parseHTML;

			exports.options = { iterations: 1, vus: 1 };

			exports.default = function() {
				var doc = parseHTML("<html><div class='something'><h1 id='top'>Lorem ipsum</h1></div></html>");

				var o = doc.find("div").get(0).attributes()

				console.log(o)
			};
