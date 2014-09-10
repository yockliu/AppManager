angular.module('app.controllers', [])

.controller('AppListCtrl', ['$scope',
  function($scope) {
    $scope.txt = 'hello this is an app list'
  }
])

.controller('CreateAppCtrl', ['$scope', '$http',
  function($scope, $http) {

    // new App.save({'name': '周末去哪儿', 'platform': 'android'}).$promise.then(function(data){
      // debugger
      // $scope.txt = data
      // console.log(data)
    // })
    $http.post('http://localhost:3000/api/app', {'name': '周末去哪儿', 'platform': 'android'})
    .success(function(data, status, headers, config) {
      console.log(data)
      $scope.txt = data
    })
    .error(function(data, status, headers, config) {
      debugger
      // called asynchronously if an error occurs
      // or server returns response with an error status.
    });
  }
])

.controller('AppDetailsCtrl', ['$scope',
  function($scope) {
    $scope.txt = 'hello this is an app details'
  }
])