'use strict';

var app = angular.module('beerz', []);

app.config(
	 function($routeProvider, $locationProvider) {
		$routeProvider.when('/beers', {
			templateUrl: '/assets/partials/beers.html',
			controller: 'BeersCtrl'
		});
		$routeProvider.when('/beers/add', {
			templateUrl: '/assets/partials/add.html',
			controller: 'AddCtrl'
		});
		$locationProvider.html5Mode(true);
	}
);
