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
      if (data.platforms.indexOf('ios') == 0)
        $scope.app.platforms = ['', 'ios']
      console.log(data)
    }).catch(function(resp) {
      console.log(resp)
    })

    $scope.update = function() {
      var app = {
        name: $scope.app.name,
        platforms: $scope.app.platforms,
        prj_path: $scope.app.prj_path
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

.controller('AppDetailsCtrl', ['$scope', '$routeParams', '$http', 'App', 'Version', 'Channel',
  function($scope, $routeParams, $http, App, Version, Channel) {
    $scope.app = new App({
      id: $routeParams.app_id
    })
    $scope.app.$get({
      app_id: $routeParams.app_id
    }).then(function(data) {
      if (data.platforms) {
        _.forEach(data.platforms, function(platform) {
          var taskUrl = '/api/build/tasks?appid=' + $routeParams.app_id + '&platform=' + platform
          $http.get(taskUrl)
            .success(function(data) {
              if (platform === 'android')
                $scope.android_tasks = data
              else if (platform === 'ios')
                $scope.ios_tasks = data
            })
            .error(function(resp) {
              console.error(resp)
            })
        })
      }
    }).catch(function(resp) {
      console.error(resp)
    })

    Version.query({
      app_id: $routeParams.app_id
    }).$promise.then(function(data) {
      $scope.versions = data
    }).catch(function(resp) {
      console.error(resp)
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

    $scope.deleteVersion = function(id, index) {
      var result = window.confirm('确定要删除吗？')
      if (result) {
        Version.delete({
          app_id: $routeParams.app_id,
          version_id: id
        }).$promise.then(function(data) {
          $scope.versions.splice(index, 1)
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

.controller('AddVersionCtrl', ['$scope', '$routeParams', 'Version',
  function($scope, $routeParams, Version) {
    $scope.platforms = angular.fromJson($routeParams.platforms)

    $scope.create = function() {
      if (!$scope.version.name) {
        alert('版本名称不能为空！')
        return
      }
      if (!$scope.version.code) {
        alert('版本 code 不能为空！')
        return
      }

      var version = new Version($scope.version)
      version.$save({
        app_id: $routeParams.app_id
      }).then(function(data) {
        console.log(data)
        location.href = '#/apps/' + $routeParams.app_id
      }).catch(function(resp) {
        console.error(resp)
      })
    }
  }
])

.controller('UpdateVersionCtrl', ['$scope', '$routeParams', 'Version',
  function($scope, $routeParams, Version) {
    $scope.isUpdate = $routeParams.isUpdate

    $scope.platforms = angular.fromJson($routeParams.platforms)

    Version.get({
      app_id: $routeParams.app_id,
      version_id: $routeParams.version_id
    }).$promise.then(function(data) {
      $scope.version = data
    }).catch(function(resp) {
      console.error(resp)
    })

    $scope.update = function() {
      var version = $scope.version
      if (!version.name) {
        alert('版本名称不能为空！')
        return
      }
      if (!version.code) {
        alert('版本 code 不能为空！')
        return
      }

      Version.update({
        app_id: $routeParams.app_id,
        version_id: $routeParams.version_id
      }, {
        code: version.code,
        name: version.name,
        platform: version.platform,
        git_index: version.git_index,
        git_tag: version.git_tag
      }).$promise.then(function(data) {
        location.href = '#/apps/' + $routeParams.app_id
      }).catch(function(resp) {
        console.error('修改版本失败！')
      })
    }
  }
])

.controller('AddChannelCtrl', ['$scope', '$routeParams', 'Channel',
  function($scope, $routeParams, Channel) {
    $scope.platforms = angular.fromJson($routeParams.platforms)
    $scope.create = function() {
      if (!$scope.channel.name) {
        alert('渠道名称不能为空！')
        return
      }
      if (!$scope.channel.code) {
        alert('渠道 code 不能为空！')
        return
      }

      var channel = new Channel($scope.channel)

      channel.$save({
        app_id: $routeParams.app_id
      }).then(function(data) {
        console.log(data)
        location.href = '#/apps/' + $routeParams.app_id
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

.controller('BuildPackagesCtrl', ['$scope', '$routeParams', '$http', 'App', 'Version', 'Channel',
  function($scope, $routeParams, $http, App, Version, Channel) {
    $scope.buildChannels = []

    $scope.app = App.get({
      app_id: $routeParams.app_id
    })
    $scope.app.$promise.then(function(app) {
      if (app && app.platforms)
        $scope.buildPlatform = app.platforms[0]
    })

    $scope.versions = Version.query({
      app_id: $routeParams.app_id
    })
    $scope.versions.$promise.then(function(versions) {
      if (versions && versions.length)
        $scope.buildVersion = versions[0].id
    })

    $scope.channels = Channel.query({
      app_id: $routeParams.app_id
    })

    $scope.build = function() {
      var data = {
        appid: $routeParams.app_id,
        platform: $scope.buildPlatform,
        versionid: $scope.buildVersion,
        channels: _.compact($scope.buildChannels)
      }

      $http.post('/api/build', data).success(function(data) {
        location.href = '#/apps/' + $routeParams.app_id
      }).error(function(resp) {
        console.error(resp)
        alert('打包失败！')
      })
    }
  }
])