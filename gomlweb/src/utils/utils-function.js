import _ from "lodash"

export function getShortImageName(imageFullName) {
  var index = imageFullName.lastIndexOf("/");
  if (index === -1)
    return imageFullName;

  return imageFullName.substring(index + 1);
}

export function addAdditionalProp(nodes) {
  for (var node of nodes) {
    // GPU 사용률 및 툴팁 설정
    for (var gpu of node.gpu_statuses) {
      if (_.isNil(gpu.memory_used_size) === true) continue;

      var percent = (gpu.memory_used_size / gpu.memory_total_size) * 100;

      gpu.mem_used_percent = Math.round(percent);
      gpu.tooltip = `total: ${gpu.memory_total_size}, used: ${gpu.memory_used_size}`;
    }

    // 컨테이너에 노드 정보 설정
    for (var container of node.containers) {
      container.nodeID = node.id;
      container.nodeAlias = node.alias;
      container.nodeHostname = node.hostname;
      container.address = node.address;
    }
  }
}

export function filterContainers(filters, containers) {
  var filterTexts = _.map(filters, "text");

  var filterd = _.filter(containers, function (container) {
    var keys = _.keys(container.labels);
    var values = _.values(container.labels);
    // keys
    var keyExist = _.isNil(_.find(keys, function (key) {
      return _.isNil(_.find(filterTexts, function (filter) {
        return key.indexOf(filter) !== -1;
      })) === false;
    })) === false;

    // value
    var valueExist = _.isNil(_.find(values, function (value) {
      return _.isNil(_.find(filterTexts, function (filter) {
        return value.indexOf(filter) !== -1;
      })) === false;
    })) === false;

    return (keyExist === true || valueExist === true);
  });

  return filterd;
}
