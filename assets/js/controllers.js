app.controller('AppCtrl', function($scope, $http, $location) {

	$scope.getUser = function(){
		$http.get('/api/user').success(function(user){
			$scope.user = user;
			console.log("Current user: ", $scope.user);
		}).error(function(e){
			console.log("Error GET'ing current user!");
		});
	}
	$scope.getUser();

	$scope.isActive = function(path) {
		console.log("path:", $location.path());
		if ($location.path() == path){
			return 'active';
		}
	}

});

app.controller('BeersCtrl', function($scope, $http) {

	$scope.beers = [
		{name:"Beer1"},
		{name:"Beer2"}
	];
	
	$http({method: 'GET', url: '/api/beers'}).
		success(function(data, status, headers, config) {
		// this callback will be called asynchronously when the response is available
			console.log("SUCCESS", status);
			console.log("data:", data);
			$scope.beers = data;
		}).
		error(function(data, status, headers, config) {
		// called asynchronously if an error occurs or server returns response with an error status.
			console.log("ERROR GET'ing /api/beers");
		});
	
	$scope.orderPredicate = 'Score';
	$scope.isReversed = true;

	$scope.setOrder = function(predicate){
		if ($scope.orderPredicate == predicate){
			$scope.isReversed = !$scope.isReversed;
		} else {
			$scope.orderPredicate = predicate;
		}		
	};

});

app.controller('AddCtrl', function($scope, $http) {


});
