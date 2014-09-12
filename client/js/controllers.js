angular.module('app.controllers', [])

.controller('AppListCtrl', ['$scope', 'App',
  function($scope, App) {
    App.query().$promise.then(function(data) {
      $scope.appList = data
    }).catch(function(resp) {
      console.log(resp)
    })

    $scope.deleteApp = function(id, index) {
      var result = window.confirm('确定要删除吗？')
      if (result) {
        App.delete({
          app_id: id
        }).$promise.then(function(data) {
          $scope.appList.splice(index, 1)
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
    $scope.create = function() {
      var app = $scope.app
      if (!app.name) {
        alert('应用名称不能为空')
        return
      }

      // remove empty item
      app.platforms = _.compact(app.platforms)

      if (!app.platforms.length) {
        alert('至少选择一个平台！')
        return
      }

      App.save(angular.toJson(app)).$promise.then(function(data) {
        location.href = '#/apps'
      }).catch(function(resp) {
        alert('创建应用失败！')
        console.log(resp)
      })
    }
  }
])

.controller('UpdateAppCtrl', ['$scope', '$routeParams', 'App',
  function($scope, $routeParams, App) {
    $scope.isUpdate = true

    App.get({
      app_id: $routeParams.app_id
    }).$promise.then(function(data) {
      $scope.app = data
      if (data.platforms.indexOf('android') == 0)
        $scope.app.platforms = ['', 'android']
      console.log(data)
    }).catch(function(resp) {
      console.log(resp)
    })

    $scope.update = function() {
      var app = {
        name: $scope.app.name,
        platforms: $scope.app.platforms
      }
      App.update({
        app_id: $scope.app.id
      }, app).$promise.then(function(data) {
        location.href = '#/apps/' + $routeParams.app_id
      }).catch(function(resp) {
        console.err(resp)
        alert('修改失败！')
      })
    }
  }
])

.controller('AppDetailsCtrl', ['$scope', '$routeParams', 'App', 'Channel',
  function($scope, $routeParams, App, Channel) {
    $scope.app = new App({
      id: $routeParams.app_id
    })
    $scope.app.$get({
      app_id: $routeParams.app_id
    }).then(function(data) {
      console.log(data)
    }).catch(function(resp) {
      console.log(resp)
    })

    Channel.query({
      app_id: $routeParams.app_id
    }).$promise.then(function(data) {
      $scope.channels = data
    }).catch(function(resp) {
      console.error(resp)
    })

    $scope.deleteApp = function(id) {
      var result = window.confirm('确定要删除吗？')
      if (result) {
        App.delete({
          app_id: id
        }).$promise.then(function(data) {
          location.href = '#/apps'
        }).catch(function(resp) {
          console.error(resp)
          alert('删除失败！')
        })
      }
    }

    $scope.deleteChannel = function(id, index) {
      var result = window.confirm('确定要删除吗？')
      if (result) {
        Channel.delete({
          app_id: $routeParams.app_id,
          channel_id: id
        }).$promise.then(function(data) {
          $scope.channels.splice(index, 1)
        }).catch(function(resp) {
          console.error(resp)
          alert('删除失败！')
        })
      }
    }
  }
])

.controller('AddChannelCtrl', ['$scope', '$routeParams', 'Channel',
  function($scope, $routeParams, Channel) {
    $scope.platforms = angular.fromJson($routeParams.platforms)
    $scope.create = function() {
      var channel = $scope.channel
      if (!channel.name) {
        alert('渠道名称不能为空！')
        return
      }
      if (!channel.code) {
        alert('渠道 code 不能为空！')
        return
      }

      var app_id = $routeParams.app_id

      var channel = new Channel($scope.channel)
      console.log(channel)
      channel.$save({
        app_id: app_id
      }).then(function(data) {
        console.log(data)
        location.href = '#/apps/' + app_id
      }).catch(function(resp) {
        alert('添加渠道失败！')
        console.error(resp)
      })
    }
  }
])

.controller('UpdateChannelCtrl', ['$scope', '$routeParams', 'Channel',
  function($scope, $routeParams, Channel) {
    $scope.isUpdate = $routeParams.isUpdate

    $scope.platforms = angular.fromJson($routeParams.platforms)

    Channel.get({
      app_id: $routeParams.app_id,
      channel_id: $routeParams.channel_id
    }).$promise.then(function(data) {
      console.log(data)
      $scope.channel = data
    }).catch(function(resp) {
      console.log(resp)
    })

    $scope.update = function() {
      var channel = $scope.channel
      if (!channel.name) {
        alert('渠道名称不能为空！')
        return
      }
      if (!channel.code) {
        alert('渠道 code 不能为空！')
        return
      }

      Channel.update({
        app_id: $routeParams.app_id,
        channel_id: $routeParams.channel_id
      }, {
        code: channel.code,
        name: channel.name,
        platform: channel.platform
      }).$promise.then(function(data) {
        location.href = '#/apps/' + $routeParams.app_id
      }).catch(function(resp) {
        console.error('修改渠道失败！')
      })
    }
  }
])