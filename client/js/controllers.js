angular.module('app.controllers', [])

.controller('AppListCtrl', ['$scope', 'App',
  function($scope, App) {
    App.query().$promise.then(function(data) {
      $scope.appList = data
    }).catch(function(resp) {
      console.log(resp)
    })

    $scope.deleteApp = function(id) {
      var result = window.confirm('确定要删除吗？')
      if (result) {
        App.delete({
          app_id: id
        }).$promise.then(function(data) {
          location.reload()
        }).catch(function(resp) {
          console.log(resp)
          alert('删除失败！')
        })
      }
    }
  }
])

.controller('CreateAppCtrl', ['$scope', 'App',
  function($scope, App) {
    App.save({
      'name': '周末去哪儿 - v2.2.1',
      'platforms': ['android']
    }).$promise.then(function(data) {
      $scope.txt = data
      console.log(data)
    }).catch(function(resp) {
      $scope.txt = resp
      console.log(resp)
    })
  }
])

.controller('AppDetailsCtrl', ['$scope', '$routeParams', 'App',
  function($scope, $routeParams, App) {
    $scope.app = new App({
      id: $routeParams.app_id
    })
    $scope.app.$get({app_id:$routeParams.app_id}).then(function(data) {
      console.log(data)
    }).catch(function(resp) {
      console.log(resp)
    })
  }
])