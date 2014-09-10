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
    $scope.submit = function() {
      var app = $scope.app
      if (!app.name) {
        alert('应用名称不能为空')
        return
      }

      if (app.platforms[0])
        app.platforms[0] = 'ios'
      if (app.platforms[1])
        app.platforms[1] = 'android'

      App.save(angular.toJson(app)).$promise.then(function(data) {
        location.href = '#/apps'
      }).catch(function(resp) {
        alert('创建应用失败！')
        console.log(resp)
      })
    }
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