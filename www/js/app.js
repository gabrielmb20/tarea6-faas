function getAll(entity) {
	fetch('https://tarea6-faas.netlify.app/api/' + entity)
	  .then((response) => response.json())
		.then((data) => {
			fetch('/template/list/' + entity + '.html')
				.then((response) => response.text())
				.then((template) => {
					var rendered = Mustache.render(template, data);
					document.getElementById('content').innerHTML = rendered;
				});
		})
}

function getById(query, entity) {
	var params = new URLSearchParams(query);
	fetch('https://tarea6-faas.netlify.app/api/' + entity + '/?id=' + params.get('id'))
	  .then((response) => response.json())
		.then((data) => {
			fetch('/template/detail/' + entity + '.html')
				.then((response) => response.text())
				.then((template) => {
					var rendered = Mustache.render(template, data);
					document.getElementById('content').innerHTML = rendered;
				});
		})
}

function home() {
	fetch('/template/home.html')
		.then((response) => response.text())
		.then((template) => {
			var rendered = Mustache.render(template, {});
			document.getElementById('content').innerHTML = rendered;
		});
}

function init() {
	router = new Navigo(null, false, '#!');
	router.on({
		'/courss': function() {
			getAll('courses');
		},
		'/groups': function() {
			getAll('groups');
		},
		'/professors': function() {
			getAll('professors');
		},
		'/courseById': function(_, query) {
			getById(query, 'courses');
		},
		'/groupById': function(_, query) {
			getById(query, 'groups');
		},
		'/professorById': function(_, query) {
			getById(query, 'professors');
		}
	});
	router.on(() => home());
	router.resolve();
}
