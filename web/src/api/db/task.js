import service from '@/utils/request'

export const listTask = (data) => {
  return service({
    url: '/db/task/list',
    method: 'post',
    data
  })
}

export const listReviewTask = (data) => {
  return service({
    url: '/db/task/review',
    method: 'post',
    data
  })
}

export const listHistoryTask = (data) => {
  return service({
    url: '/db/task/history',
    method: 'post',
    data
  })
}
export const getTask = (data) => {
  return service({
    url: '/db/task?id=' + data,
    method: 'get',
  })
}

export const updateTask = (data) => {
  return service({
    url: '/db/task',
    method: 'put',
    data
  })
}

export const createTask = (data) => {
  return service({
    url: '/db/task',
    method: 'post',
    data
  })
}

export const listBackup = (data) => {
  return service({
    url: '/db/backup/list',
    method: 'post',
    data
  })
}

export const rollback = (data) => {
  return service({
    url: '/db/backup/rollback',
    method: 'post',
    data
  })
}
