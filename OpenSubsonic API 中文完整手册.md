# OpenSubsonic API 中文完整手册

## 1. 认证与基础说明

### API Base URL 说明
API 请求的基础 URL 格式为 `http://server/rest/`，其中 `server` 是您的 OpenSubsonic 服务器地址。

### 认证方式
OpenSubsonic API 支持多种认证方式，包括：
- **Token 认证**：使用预生成的 API token
- **Salt/MD5 认证**：使用用户名、时间戳、盐值和密码的 MD5 哈希

### 通用请求参数

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本（如 1.16.1） |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式（xml 或 json） |

### 版本兼容说明
OpenSubsonic API 从版本 1.8.0 开始提供了基于 ID3 标签的媒体集合访问方法，而不是基于文件结构。例如，使用 ID3 标签浏览集合应使用 getArtists、getArtist 和 getAlbum 方法，而使用文件结构浏览则应使用 getIndexes 和 getMusicDirectory。

---

## 2. API 分类目录

### 系统 API
- ping
- getLicense
- getOpenSubsonicExtensions
- tokenInfo

### 浏览 API
- getMusicFolders
- getIndexes
- getMusicDirectory
- getGenres
- getArtists
- getArtist
- getAlbum
- getSong
- getVideos
- getVideoInfo
- getArtistInfo
- getArtistInfo2
- getAlbumInfo
- getAlbumInfo2
- getSimilarSongs
- getSimilarSongs2
- getTopSongs

### 专辑/歌曲列表 API
- getAlbumList
- getAlbumList2
- getRandomSongs
- getSongsByGenre
- getNowPlaying
- getStarred
- getStarred2

### 搜索 API
- search
- search2
- search3

### 播放列表 API
- getPlaylists
- getPlaylist
- createPlaylist
- updatePlaylist
- deletePlaylist

### 媒体检索 API
- stream
- download
- hls
- getCaptions
- getCoverArt
- getLyrics
- getAvatar
- getLyricsBySongId

### 媒体标注 API
- star
- unstar
- setRating
- scrobble

### 共享 API
- getShares
- createShare
- updateShare
- deleteShare

### 播客 API
- getPodcasts
- getNewestPodcasts
- refreshPodcasts
- createPodcastChannel
- deletePodcastChannel
- deletePodcastEpisode
- downloadPodcastEpisode

### 点唱机 API
- jukeboxControl

### 网络广播 API
- getInternetRadioStations
- createInternetRadioStation
- updateInternetRadioStation
- deleteInternetRadioStation

### 聊天 API
- getChatMessages
- addChatMessage

### 用户管理 API
- getUser
- getUsers
- createUser
- updateUser
- deleteUser
- changePassword

### 书签 API
- getBookmarks
- createBookmark
- deleteBookmark
- getPlayQueue
- savePlayQueue

### 媒体库扫描 API
- getScanStatus
- startScan

---

## 3. 每个 API 的详细说明

### 3.1 ping

**接口路径：**
`/rest/ping.view`

**HTTP 方法：**
GET

**功能说明：**
用于测试与服务器的连接性。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |

**请求示例：**

```
http://server/rest/ping.view?u=admin&t=token&s=salt&v=1.16.1&c=client
```

**成功返回示例（JSON）：**

```json
{
  "subsonic-response": {
    "status": "ok",
    "version": "1.16.1",
    "type": "OpenSubsonic",
    "serverVersion": "1.0.0",
    "openSubsonic": true
  }
}
```

**成功返回示例（XML）：**

```xml
<subsonic-response xmlns="http://subsonic.org/restapi" status="ok" version="1.16.1" type="OpenSubsonic" serverVersion="1.0.0" openSubsonic="true">
</subsonic-response>
```

### 3.2 getLicense

**接口路径：**
`/rest/getLicense.view`

**HTTP 方法：**
GET

**功能说明：**
获取有关软件许可证的详细信息。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |

**请求示例：**

```
http://server/rest/getLicense.view?u=admin&t=token&s=salt&v=1.16.1&c=client
```

### 3.3 getOpenSubsonicExtensions

**接口路径：**
`/rest/getOpenSubsonicExtensions.view`

**HTTP 方法：**
GET

**功能说明：**
列出此服务器支持的 OpenSubsonic 扩展。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |

**请求示例：**

```
http://server/rest/getOpenSubsonicExtensions.view?u=admin&t=token&s=salt&v=1.16.1&c=client
```

### 3.4 tokenInfo

**接口路径：**
`/rest/tokenInfo.view`

**HTTP 方法：**
GET

**功能说明：**
获取令牌信息。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |

**请求示例：**

```
http://server/rest/tokenInfo.view?u=admin&t=token&s=salt&v=1.16.1&c=client
```

### 3.5 getMusicFolders

**接口路径：**
`/rest/getMusicFolders.view`

**HTTP 方法：**
GET

**功能说明：**
返回所有配置的顶级音乐文件夹。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |

**请求示例：**

```
http://server/rest/getMusicFolders.view?u=admin&t=token&s=salt&v=1.16.1&c=client
```

### 3.6 getIndexes

**接口路径：**
`/rest/getIndexes.view`

**HTTP 方法：**
GET

**功能说明：**
返回所有艺术家的索引结构。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| musicFolderId | 字符串 | 否 | 无 | 音乐文件夹 ID |

**请求示例：**

```
http://server/rest/getIndexes.view?u=admin&t=token&s=salt&v=1.16.1&c=client
```

### 3.7 getMusicDirectory

**接口路径：**
`/rest/getMusicDirectory.view`

**HTTP 方法：**
GET

**功能说明：**
返回音乐目录中所有文件的列表。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 目录 ID |

**请求示例：**

```
http://server/rest/getMusicDirectory.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1
```

### 3.8 getGenres

**接口路径：**
`/rest/getGenres.view`

**HTTP 方法：**
GET

**功能说明：**
返回所有流派。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |

**请求示例：**

```
http://server/rest/getGenres.view?u=admin&t=token&s=salt&v=1.16.1&c=client
```

### 3.9 getArtists

**接口路径：**
`/rest/getArtists.view`

**HTTP 方法：**
GET

**功能说明：**
返回所有艺术家。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| musicFolderId | 字符串 | 否 | 无 | 音乐文件夹 ID |

**请求示例：**

```
http://server/rest/getArtists.view?u=admin&t=token&s=salt&v=1.16.1&c=client
```

### 3.10 getArtist

**接口路径：**
`/rest/getArtist.view`

**HTTP 方法：**
GET

**功能说明：**
返回艺术家的详细信息。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 艺术家 ID |

**请求示例：**

```
http://server/rest/getArtist.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1
```

### 3.11 getAlbum

**接口路径：**
`/rest/getAlbum.view`

**HTTP 方法：**
GET

**功能说明：**
返回专辑的详细信息。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 专辑 ID |

**请求示例：**

```
http://server/rest/getAlbum.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1
```

### 3.12 getSong

**接口路径：**
`/rest/getSong.view`

**HTTP 方法：**
GET

**功能说明：**
返回歌曲的详细信息。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 歌曲 ID |

**请求示例：**

```
http://server/rest/getSong.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1
```

### 3.13 getVideos

**接口路径：**
`/rest/getVideos.view`

**HTTP 方法：**
GET

**功能说明：**
返回所有视频文件。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |

**请求示例：**

```
http://server/rest/getVideos.view?u=admin&t=token&s=salt&v=1.16.1&c=client
```

### 3.14 getVideoInfo

**接口路径：**
`/rest/getVideoInfo.view`

**HTTP 方法：**
GET

**功能说明：**
返回视频的详细信息。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 视频 ID |

**请求示例：**

```
http://server/rest/getVideoInfo.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1
```

### 3.15 getArtistInfo

**接口路径：**
`/rest/getArtistInfo.view`

**HTTP 方法：**
GET

**功能说明：**
返回艺术家的详细信息。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 艺术家 ID |

**请求示例：**

```
http://server/rest/getArtistInfo.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1
```

### 3.16 getArtistInfo2

**接口路径：**
`/rest/getArtistInfo2.view`

**HTTP 方法：**
GET

**功能说明：**
返回艺术家的详细信息（增强版）。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 艺术家 ID |

**请求示例：**

```
http://server/rest/getArtistInfo2.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1
```

### 3.17 getAlbumInfo

**接口路径：**
`/rest/getAlbumInfo.view`

**HTTP 方法：**
GET

**功能说明：**
返回专辑的详细信息。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 专辑 ID |

**请求示例：**

```
http://server/rest/getAlbumInfo.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1
```

### 3.18 getAlbumInfo2

**接口路径：**
`/rest/getAlbumInfo2.view`

**HTTP 方法：**
GET

**功能说明：**
返回专辑的详细信息（增强版）。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 专辑 ID |

**请求示例：**

```
http://server/rest/getAlbumInfo2.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1
```

### 3.19 getSimilarSongs

**接口路径：**
`/rest/getSimilarSongs.view`

**HTTP 方法：**
GET

**功能说明：**
返回给定艺术家和类似艺术家的随机歌曲集合。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 歌曲或艺术家 ID |
| count | 整数 | 否 | 50 | 返回的歌曲数量 |

**请求示例：**

```
http://server/rest/getSimilarSongs.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1&count=10
```

### 3.20 getSimilarSongs2

**接口路径：**
`/rest/getSimilarSongs2.view`

**HTTP 方法：**
GET

**功能说明：**
返回给定艺术家和类似艺术家的随机歌曲集合（增强版）。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 歌曲或艺术家 ID |
| count | 整数 | 否 | 50 | 返回的歌曲数量 |

**请求示例：**

```
http://server/rest/getSimilarSongs2.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1&count=10
```

### 3.21 getTopSongs

**接口路径：**
`/rest/getTopSongs.view`

**HTTP 方法：**
GET

**功能说明：**
返回给定艺术家的热门歌曲。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 艺术家 ID |
| count | 整数 | 否 | 50 | 返回的歌曲数量 |

**请求示例：**

```
http://server/rest/getTopSongs.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1&count=10
```

### 3.22 getAlbumList

**接口路径：**
`/rest/getAlbumList.view`

**HTTP 方法：**
GET

**功能说明：**
返回随机、最新、评分最高等类型的专辑列表。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| type | 字符串 | 是 | 无 | 列表类型（random, newest, highest, frequent, recent） |
| size | 整数 | 否 | 10 | 返回的专辑数量 |
| offset | 整数 | 否 | 0 | 偏移量 |
| musicFolderId | 字符串 | 否 | 无 | 音乐文件夹 ID |

**请求示例：**

```
http://server/rest/getAlbumList.view?u=admin&t=token&s=salt&v=1.16.1&c=client&type=random&size=10
```

### 3.23 getAlbumList2

**接口路径：**
`/rest/getAlbumList2.view`

**HTTP 方法：**
GET

**功能说明：**
返回随机、最新、评分最高等类型的专辑列表（增强版）。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| type | 字符串 | 是 | 无 | 列表类型 |
| size | 整数 | 否 | 10 | 返回的专辑数量 |
| offset | 整数 | 否 | 0 | 偏移量 |
| musicFolderId | 字符串 | 否 | 无 | 音乐文件夹 ID |

**请求示例：**

```
http://server/rest/getAlbumList2.view?u=admin&t=token&s=salt&v=1.16.1&c=client&type=random&size=10
```

### 3.24 getRandomSongs

**接口路径：**
`/rest/getRandomSongs.view`

**HTTP 方法：**
GET

**功能说明：**
返回符合给定条件的随机歌曲。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| size | 整数 | 否 | 50 | 返回的歌曲数量 |
| genre | 字符串 | 否 | 无 | 流派过滤 |
| fromYear | 整数 | 否 | 无 | 起始年份 |
| toYear | 整数 | 否 | 无 | 结束年份 |
| musicFolderId | 字符串 | 否 | 无 | 音乐文件夹 ID |

**请求示例：**

```
http://server/rest/getRandomSongs.view?u=admin&t=token&s=salt&v=1.16.1&c=client&size=10
```

### 3.25 getSongsByGenre

**接口路径：**
`/rest/getSongsByGenre.view`

**HTTP 方法：**
GET

**功能说明：**
返回给定流派的歌曲。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| genre | 字符串 | 是 | 无 | 流派名称 |
| count | 整数 | 否 | 50 | 返回的歌曲数量 |
| offset | 整数 | 否 | 0 | 偏移量 |

**请求示例：**

```
http://server/rest/getSongsByGenre.view?u=admin&t=token&s=salt&v=1.16.1&c=client&genre=Rock&count=10
```

### 3.26 getNowPlaying

**接口路径：**
`/rest/getNowPlaying.view`

**HTTP 方法：**
GET

**功能说明：**
返回所有用户当前正在播放的内容。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |

**请求示例：**

```
http://server/rest/getNowPlaying.view?u=admin&t=token&s=salt&v=1.16.1&c=client
```

### 3.27 getStarred

**接口路径：**
`/rest/getStarred.view`

**HTTP 方法：**
GET

**功能说明：**
返回已加星标的歌曲、专辑和艺术家。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| musicFolderId | 字符串 | 否 | 无 | 音乐文件夹 ID |

**请求示例：**

```
http://server/rest/getStarred.view?u=admin&t=token&s=salt&v=1.16.1&c=client
```

### 3.28 getStarred2

**接口路径：**
`/rest/getStarred2.view`

**HTTP 方法：**
GET

**功能说明：**
返回已加星标的歌曲、专辑和艺术家（增强版）。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| musicFolderId | 字符串 | 否 | 无 | 音乐文件夹 ID |

**请求示例：**

```
http://server/rest/getStarred2.view?u=admin&t=token&s=salt&v=1.16.1&c=client
```

### 3.29 search

**接口路径：**
`/rest/search.view`

**HTTP 方法：**
GET

**功能说明：**
搜索并返回匹配给定搜索条件的文件。支持分页。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| query | 字符串 | 是 | 无 | 搜索查询 |
| artistCount | 整数 | 否 | 20 | 返回的艺术家数量 |
| albumCount | 整数 | 否 | 20 | 返回的专辑数量 |
| songCount | 整数 | 否 | 20 | 返回的歌曲数量 |

**请求示例：**

```
http://server/rest/search.view?u=admin&t=token&s=salt&v=1.16.1&c=client&query=Beatles
```

### 3.30 search2

**接口路径：**
`/rest/search2.view`

**HTTP 方法：**
GET

**功能说明：**
搜索并返回匹配给定搜索条件的文件（增强版）。支持分页。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| query | 字符串 | 是 | 无 | 搜索查询 |
| artistCount | 整数 | 否 | 20 | 返回的艺术家数量 |
| albumCount | 整数 | 否 | 20 | 返回的专辑数量 |
| songCount | 整数 | 否 | 20 | 返回的歌曲数量 |

**请求示例：**

```
http://server/rest/search2.view?u=admin&t=token&s=salt&v=1.16.1&c=client&query=Beatles
```

### 3.31 search3

**接口路径：**
`/rest/search3.view`

**HTTP 方法：**
GET

**功能说明：**
搜索并返回匹配给定搜索条件的专辑、艺术家和歌曲。支持分页。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| query | 字符串 | 是 | 无 | 搜索查询 |
| artistCount | 整数 | 否 | 20 | 返回的艺术家数量 |
| albumCount | 整数 | 否 | 20 | 返回的专辑数量 |
| songCount | 整数 | 否 | 20 | 返回的歌曲数量 |
| offset | 整数 | 否 | 0 | 偏移量 |
| all | 布尔值 | 否 | false | 是否搜索所有字段 |

**请求示例：**

```
http://server/rest/search3.view?u=admin&t=token&s=salt&v=1.16.1&c=client&query=Beatles
```

### 3.32 getPlaylists

**接口路径：**
`/rest/getPlaylists.view`

**HTTP 方法：**
GET

**功能说明：**
返回用户有权播放的所有播放列表。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |

**请求示例：**

```
http://server/rest/getPlaylists.view?u=admin&t=token&s=salt&v=1.16.1&c=client
```

### 3.33 getPlaylist

**接口路径：**
`/rest/getPlaylist.view`

**HTTP 方法：**
GET

**功能说明：**
返回已保存播放列表中的文件列表。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 播放列表 ID |

**请求示例：**

```
http://server/rest/getPlaylist.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1
```

### 3.34 createPlaylist

**接口路径：**
`/rest/createPlaylist.view`

**HTTP 方法：**
GET

**功能说明：**
创建（或更新）播放列表。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 否 | 无 | 播放列表 ID（更新现有播放列表时使用） |
| name | 字符串 | 是 | 无 | 播放列表名称 |
| comment | 字符串 | 否 | 无 | 播放列表注释 |
| songId | 字符串 | 否 | 无 | 歌曲 ID 列表（逗号分隔） |

**请求示例：**

```
http://server/rest/createPlaylist.view?u=admin&t=token&s=salt&v=1.16.1&c=client&name=My Playlist
```

### 3.35 updatePlaylist

**接口路径：**
`/rest/updatePlaylist.view`

**HTTP 方法：**
GET

**功能说明：**
更新播放列表。只有播放列表的所有者可以更新它。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 播放列表 ID |
| name | 字符串 | 否 | 无 | 播放列表名称 |
| comment | 字符串 | 否 | 无 | 播放列表注释 |
| public | 布尔值 | 否 | false | 是否公开 |
| songIdToAdd | 字符串 | 否 | 无 | 要添加的歌曲 ID 列表（逗号分隔） |
| songIndexToRemove | 字符串 | 否 | 无 | 要删除的歌曲索引列表（逗号分隔） |

**请求示例：**

```
http://server/rest/updatePlaylist.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1&name=Updated Playlist
```

### 3.36 deletePlaylist

**接口路径：**
`/rest/deletePlaylist.view`

**HTTP 方法：**
GET

**功能说明：**
删除已保存的播放列表。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 播放列表 ID |

**请求示例：**

```
http://server/rest/deletePlaylist.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1
```

### 3.37 stream

**接口路径：**
`/rest/stream.view`

**HTTP 方法：**
GET

**功能说明：**
流式传输给定的媒体文件。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| id | 字符串 | 是 | 无 | 媒体文件 ID |
| maxBitRate | 整数 | 否 | 无 | 最大比特率 |
| format | 字符串 | 否 | 无 | 输出格式 |
| timeOffset | 整数 | 否 | 0 | 时间偏移（秒） |
| size | 整数 | 否 | 无 | 要返回的数据大小 |
| estimateContentLength | 布尔值 | 否 | false | 是否估计内容长度 |

**请求示例：**

```
http://server/rest/stream.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1
```

### 3.38 download

**接口路径：**
`/rest/download.view`

**HTTP 方法：**
GET

**功能说明：**
下载给定的媒体文件。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 媒体文件 ID |

**请求示例：**

```
http://server/rest/download.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1
```

### 3.39 hls

**接口路径：**
`/rest/hls.view`

**HTTP 方法：**
GET

**功能说明：**
下载给定的媒体文件（HLS 格式）。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| id | 字符串 | 是 | 无 | 媒体文件 ID |
| maxBitRate | 整数 | 否 | 无 | 最大比特率 |

**请求示例：**

```
http://server/rest/hls.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1
```

### 3.40 getCaptions

**接口路径：**
`/rest/getCaptions.view`

**HTTP 方法：**
GET

**功能说明：**
返回视频的字幕。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 视频 ID |

**请求示例：**

```
http://server/rest/getCaptions.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1
```

### 3.41 getCoverArt

**接口路径：**
`/rest/getCoverArt.view`

**HTTP 方法：**
GET

**功能说明：**
返回封面艺术图像。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| id | 字符串 | 是 | 无 | 封面艺术 ID |
| size | 整数 | 否 | 无 | 图像大小 |

**请求示例：**

```
http://server/rest/getCoverArt.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1&size=300
```

### 3.42 getLyrics

**接口路径：**
`/rest/getLyrics.view`

**HTTP 方法：**
GET

**功能说明：**
搜索并返回给定歌曲的歌词。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| artist | 字符串 | 是 | 无 | 艺术家名称 |
| title | 字符串 | 是 | 无 | 歌曲标题 |

**请求示例：**

```
http://server/rest/getLyrics.view?u=admin&t=token&s=salt&v=1.16.1&c=client&artist=Beatles&title=Hey Jude
```

### 3.43 getAvatar

**接口路径：**
`/rest/getAvatar.view`

**HTTP 方法：**
GET

**功能说明：**
返回用户的头像（个人图像）。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| username | 字符串 | 是 | 无 | 用户名 |

**请求示例：**

```
http://server/rest/getAvatar.view?u=admin&t=token&s=salt&v=1.16.1&c=client&username=admin
```

### 3.44 getLyricsBySongId

**接口路径：**
`/rest/getLyricsBySongId.view`

**HTTP 方法：**
GET

**功能说明：**
通过歌曲 ID 获取歌词，支持同步歌词、多种语言。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 歌曲 ID |

**请求示例：**

```
http://server/rest/getLyricsBySongId.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1
```

### 3.45 star

**接口路径：**
`/rest/star.view`

**HTTP 方法：**
GET

**功能说明：**
为歌曲、专辑或艺术家添加星标。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 媒体项 ID |
| albumId | 字符串 | 否 | 无 | 专辑 ID |
| artistId | 字符串 | 否 | 无 | 艺术家 ID |

**请求示例：**

```
http://server/rest/star.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1
```

### 3.46 unstar

**接口路径：**
`/rest/unstar.view`

**HTTP 方法：**
GET

**功能说明：**
移除歌曲、专辑或艺术家的星标。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 媒体项 ID |
| albumId | 字符串 | 否 | 无 | 专辑 ID |
| artistId | 字符串 | 否 | 无 | 艺术家 ID |

**请求示例：**

```
http://server/rest/unstar.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1
```

### 3.47 setRating

**接口路径：**
`/rest/setRating.view`

**HTTP 方法：**
GET

**功能说明：**
设置音乐文件的评分。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 媒体文件 ID |
| rating | 整数 | 是 | 无 | 评分（1-5，0 表示移除评分） |

**请求示例：**

```
http://server/rest/setRating.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1&rating=5
```

### 3.48 scrobble

**接口路径：**
`/rest/scrobble.view`

**HTTP 方法：**
GET

**功能说明：**
注册一个或多个媒体文件的本地播放。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 媒体文件 ID |
| time | 整数 | 否 | 无 | 播放时间戳 |
| submission | 布尔值 | 否 | true | 是否提交到 Last.fm |

**请求示例：**

```
http://server/rest/scrobble.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1
```

### 3.49 getShares

**接口路径：**
`/rest/getShares.view`

**HTTP 方法：**
GET

**功能说明：**
返回用户有权管理的共享媒体信息。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |

**请求示例：**

```
http://server/rest/getShares.view?u=admin&t=token&s=salt&v=1.16.1&c=client
```

### 3.50 createShare

**接口路径：**
`/rest/createShare.view`

**HTTP 方法：**
GET

**功能说明：**
创建一个公共 URL，任何人都可以使用它来流式传输服务器上的音乐或视频。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 要共享的媒体文件 ID（多个 ID 用逗号分隔） |
| description | 字符串 | 否 | 无 | 共享描述 |
| expires | 整数 | 否 | 无 | 过期时间（毫秒） |

**请求示例：**

```
http://server/rest/createShare.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1,2,3
```

### 3.51 updateShare

**接口路径：**
`/rest/updateShare.view`

**HTTP 方法：**
GET

**功能说明：**
更新现有共享的描述和/或过期日期。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 共享 ID |
| description | 字符串 | 否 | 无 | 共享描述 |
| expires | 整数 | 否 | 无 | 过期时间（毫秒） |

**请求示例：**

```
http://server/rest/updateShare.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1&description=Updated Share
```

### 3.52 deleteShare

**接口路径：**
`/rest/deleteShare.view`

**HTTP 方法：**
GET

**功能说明：**
删除现有共享。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 共享 ID |

**请求示例：**

```
http://server/rest/deleteShare.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1
```

### 3.53 getPodcasts

**接口路径：**
`/rest/getPodcasts.view`

**HTTP 方法：**
GET

**功能说明：**
返回服务器订阅的所有播客频道，以及（可选）它们的剧集。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| includeEpisodes | 布尔值 | 否 | false | 是否包含剧集 |

**请求示例：**

```
http://server/rest/getPodcasts.view?u=admin&t=token&s=salt&v=1.16.1&c=client&includeEpisodes=true
```

### 3.54 getNewestPodcasts

**接口路径：**
`/rest/getNewestPodcasts.view`

**HTTP 方法：**
GET

**功能说明：**
返回最近发布的播客剧集。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| count | 整数 | 否 | 20 | 返回的剧集数量 |

**请求示例：**

```
http://server/rest/getNewestPodcasts.view?u=admin&t=token&s=salt&v=1.16.1&c=client&count=10
```

### 3.55 refreshPodcasts

**接口路径：**
`/rest/refreshPodcasts.view`

**HTTP 方法：**
GET

**功能说明：**
请求服务器检查新的播客剧集。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |

**请求示例：**

```
http://server/rest/refreshPodcasts.view?u=admin&t=token&s=salt&v=1.16.1&c=client
```

### 3.56 createPodcastChannel

**接口路径：**
`/rest/createPodcastChannel.view`

**HTTP 方法：**
GET

**功能说明：**
添加新的播客频道。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| url | 字符串 | 是 | 无 | 播客频道 URL |

**请求示例：**

```
http://server/rest/createPodcastChannel.view?u=admin&t=token&s=salt&v=1.16.1&c=client&url=http://example.com/podcast.rss
```

### 3.57 deletePodcastChannel

**接口路径：**
`/rest/deletePodcastChannel.view`

**HTTP 方法：**
GET

**功能说明：**
删除播客频道。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 播客频道 ID |

**请求示例：**

```
http://server/rest/deletePodcastChannel.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1
```

### 3.58 deletePodcastEpisode

**接口路径：**
`/rest/deletePodcastEpisode.view`

**HTTP 方法：**
GET

**功能说明：**
删除播客剧集。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 播客剧集 ID |

**请求示例：**

```
http://server/rest/deletePodcastEpisode.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1
```

### 3.59 downloadPodcastEpisode

**接口路径：**
`/rest/downloadPodcastEpisode.view`

**HTTP 方法：**
GET

**功能说明：**
请求服务器开始下载给定的播客剧集。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 播客剧集 ID |

**请求示例：**

```
http://server/rest/downloadPodcastEpisode.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1
```

### 3.60 jukeboxControl

**接口路径：**
`/rest/jukeboxControl.view`

**HTTP 方法：**
GET

**功能说明：**
控制点唱机，即直接在服务器的音频硬件上播放。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| action | 字符串 | 是 | 无 | 操作（start, stop, pause, resume, skip, setGain） |
| index | 整数 | 否 | 无 | 播放队列索引 |
| offset | 整数 | 否 | 无 | 时间偏移（秒） |
| gain | 浮点数 | 否 | 无 | 音量增益（0.0-1.0） |

**请求示例：**

```
http://server/rest/jukeboxControl.view?u=admin&t=token&s=salt&v=1.16.1&c=client&action=start
```

### 3.61 getInternetRadioStations

**接口路径：**
`/rest/getInternetRadioStations.view`

**HTTP 方法：**
GET

**功能说明：**
返回所有网络广播电台。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |

**请求示例：**

```
http://server/rest/getInternetRadioStations.view?u=admin&t=token&s=salt&v=1.16.1&c=client
```

### 3.62 createInternetRadioStation

**接口路径：**
`/rest/createInternetRadioStation.view`

**HTTP 方法：**
GET

**功能说明：**
添加新的网络广播电台。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| name | 字符串 | 是 | 无 | 电台名称 |
| streamUrl | 字符串 | 是 | 无 | 流 URL |
| homepageUrl | 字符串 | 否 | 无 | 主页 URL |

**请求示例：**

```
http://server/rest/createInternetRadioStation.view?u=admin&t=token&s=salt&v=1.16.1&c=client&name=Radio Station&streamUrl=http://example.com/stream
```

### 3.63 updateInternetRadioStation

**接口路径：**
`/rest/updateInternetRadioStation.view`

**HTTP 方法：**
GET

**功能说明：**
更新现有的网络广播电台。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 电台 ID |
| name | 字符串 | 否 | 无 | 电台名称 |
| streamUrl | 字符串 | 否 | 无 | 流 URL |
| homepageUrl | 字符串 | 否 | 无 | 主页 URL |

**请求示例：**

```
http://server/rest/updateInternetRadioStation.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1&name=Updated Radio
```

### 3.64 deleteInternetRadioStation

**接口路径：**
`/rest/deleteInternetRadioStation.view`

**HTTP 方法：**
GET

**功能说明：**
删除现有的网络广播电台。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 电台 ID |

**请求示例：**

```
http://server/rest/deleteInternetRadioStation.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1
```

### 3.65 getChatMessages

**接口路径：**
`/rest/getChatMessages.view`

**HTTP 方法：**
GET

**功能说明：**
返回当前可见（未过期）的聊天消息。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| since | 整数 | 否 | 无 | 只返回此时间戳之后的消息 |

**请求示例：**

```
http://server/rest/getChatMessages.view?u=admin&t=token&s=salt&v=1.16.1&c=client
```

### 3.66 addChatMessage

**接口路径：**
`/rest/addChatMessage.view`

**HTTP 方法：**
GET

**功能说明：**
向聊天日志添加消息。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| message | 字符串 | 是 | 无 | 消息内容 |

**请求示例：**

```
http://server/rest/addChatMessage.view?u=admin&t=token&s=salt&v=1.16.1&c=client&message=Hello
```

### 3.67 getUser

**接口路径：**
`/rest/getUser.view`

**HTTP 方法：**
GET

**功能说明：**
获取有关给定用户的详细信息，包括其授权角色和文件夹访问权限。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| username | 字符串 | 是 | 无 | 用户名 |

**请求示例：**

```
http://server/rest/getUser.view?u=admin&t=token&s=salt&v=1.16.1&c=client&username=user
```

### 3.68 getUsers

**接口路径：**
`/rest/getUsers.view`

**HTTP 方法：**
GET

**功能说明：**
获取所有用户的详细信息，包括他们的授权角色和文件夹访问权限。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |

**请求示例：**

```
http://server/rest/getUsers.view?u=admin&t=token&s=salt&v=1.16.1&c=client
```

### 3.69 createUser

**接口路径：**
`/rest/createUser.view`

**HTTP 方法：**
GET

**功能说明：**
在服务器上创建新用户。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| username | 字符串 | 是 | 无 | 用户名 |
| password | 字符串 | 是 | 无 | 密码 |
| email | 字符串 | 否 | 无 | 电子邮件 |
| ldapAuthenticated | 布尔值 | 否 | false | 是否通过 LDAP 认证 |
| adminRole | 布尔值 | 否 | false | 是否具有管理员角色 |
| settingsRole | 布尔值 | 否 | false | 是否具有设置角色 |
| streamRole | 布尔值 | 否 | true | 是否具有流角色 |
| jukeboxRole | 布尔值 | 否 | false | 是否具有点唱机角色 |
| downloadRole | 布尔值 | 否 | false | 是否具有下载角色 |
| uploadRole | 布尔值 | 否 | false | 是否具有上传角色 |
| playlistRole | 布尔值 | 否 | false | 是否具有播放列表角色 |
| coverArtRole | 布尔值 | 否 | false | 是否具有封面艺术角色 |
| commentRole | 布尔值 | 否 | false | 是否具有评论角色 |
| podcastRole | 布尔值 | 否 | false | 是否具有播客角色 |
| shareRole | 布尔值 | 否 | false | 是否具有共享角色 |
| musicFolderId | 字符串 | 否 | 无 | 音乐文件夹 ID 列表（逗号分隔） |

**请求示例：**

```
http://server/rest/createUser.view?u=admin&t=token&s=salt&v=1.16.1&c=client&username=newuser&password=password
```

### 3.70 updateUser

**接口路径：**
`/rest/updateUser.view`

**HTTP 方法：**
GET

**功能说明：**
修改服务器上的现有用户。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| username | 字符串 | 是 | 无 | 用户名 |
| password | 字符串 | 否 | 无 | 密码 |
| email | 字符串 | 否 | 无 | 电子邮件 |
| ldapAuthenticated | 布尔值 | 否 | false | 是否通过 LDAP 认证 |
| adminRole | 布尔值 | 否 | false | 是否具有管理员角色 |
| settingsRole | 布尔值 | 否 | false | 是否具有设置角色 |
| streamRole | 布尔值 | 否 | true | 是否具有流角色 |
| jukeboxRole | 布尔值 | 否 | false | 是否具有点唱机角色 |
| downloadRole | 布尔值 | 否 | false | 是否具有下载角色 |
| uploadRole | 布尔值 | 否 | false | 是否具有上传角色 |
| playlistRole | 布尔值 | 否 | false | 是否具有播放列表角色 |
| coverArtRole | 布尔值 | 否 | false | 是否具有封面艺术角色 |
| commentRole | 布尔值 | 否 | false | 是否具有评论角色 |
| podcastRole | 布尔值 | 否 | false | 是否具有播客角色 |
| shareRole | 布尔值 | 否 | false | 是否具有共享角色 |
| musicFolderId | 字符串 | 否 | 无 | 音乐文件夹 ID 列表（逗号分隔） |

**请求示例：**

```
http://server/rest/updateUser.view?u=admin&t=token&s=salt&v=1.16.1&c=client&username=user&email=user@example.com
```

### 3.71 deleteUser

**接口路径：**
`/rest/deleteUser.view`

**HTTP 方法：**
GET

**功能说明：**
在服务器上删除现有用户。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| username | 字符串 | 是 | 无 | 用户名 |

**请求示例：**

```
http://server/rest/deleteUser.view?u=admin&t=token&s=salt&v=1.16.1&c=client&username=user
```

### 3.72 changePassword

**接口路径：**
`/rest/changePassword.view`

**HTTP 方法：**
GET

**功能说明：**
更改服务器上现有用户的密码。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| username | 字符串 | 是 | 无 | 用户名 |
| newPassword | 字符串 | 是 | 无 | 新密码 |

**请求示例：**

```
http://server/rest/changePassword.view?u=admin&t=token&s=salt&v=1.16.1&c=client&username=user&newPassword=newpassword
```

### 3.73 getBookmarks

**接口路径：**
`/rest/getBookmarks.view`

**HTTP 方法：**
GET

**功能说明：**
返回此用户的所有书签。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |

**请求示例：**

```
http://server/rest/getBookmarks.view?u=admin&t=token&s=salt&v=1.16.1&c=client
```

### 3.74 createBookmark

**接口路径：**
`/rest/createBookmark.view`

**HTTP 方法：**
GET

**功能说明：**
创建或更新书签。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 媒体文件 ID |
| position | 整数 | 是 | 无 | 书签位置（秒） |
| comment | 字符串 | 否 | 无 | 书签注释 |

**请求示例：**

```
http://server/rest/createBookmark.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1&position=120
```

### 3.75 deleteBookmark

**接口路径：**
`/rest/deleteBookmark.view`

**HTTP 方法：**
GET

**功能说明：**
删除书签。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 书签 ID |

**请求示例：**

```
http://server/rest/deleteBookmark.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1
```

### 3.76 getPlayQueue

**接口路径：**
`/rest/getPlayQueue.view`

**HTTP 方法：**
GET

**功能说明：**
返回此用户的播放队列状态。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |

**请求示例：**

```
http://server/rest/getPlayQueue.view?u=admin&t=token&s=salt&v=1.16.1&c=client
```

### 3.77 savePlayQueue

**接口路径：**
`/rest/savePlayQueue.view`

**HTTP 方法：**
GET

**功能说明：**
保存此用户的播放队列状态。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |
| id | 字符串 | 是 | 无 | 要添加到播放队列的歌曲 ID 列表（逗号分隔） |
| current | 字符串 | 否 | 无 | 当前播放的歌曲 ID |
| position | 整数 | 否 | 0 | 当前播放位置（秒） |

**请求示例：**

```
http://server/rest/savePlayQueue.view?u=admin&t=token&s=salt&v=1.16.1&c=client&id=1,2,3
```

### 3.78 getScanStatus

**接口路径：**
`/rest/getScanStatus.view`

**HTTP 方法：**
GET

**功能说明：**
返回媒体库扫描的当前状态。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |

**请求示例：**

```
http://server/rest/getScanStatus.view?u=admin&t=token&s=salt&v=1.16.1&c=client
```

### 3.79 startScan

**接口路径：**
`/rest/startScan.view`

**HTTP 方法：**
GET

**功能说明：**
启动媒体库的重新扫描。

**请求参数：**

| 参数名 | 类型 | 是否必填 | 默认值 | 说明 |
|--------|------|----------|--------|------|
| u | 字符串 | 是 | 无 | 用户名 |
| t | 字符串 | 是 | 无 | 时间戳或认证令牌 |
| s | 字符串 | 是 | 无 | 盐值（用于 MD5 认证） |
| v | 字符串 | 是 | 无 | API 版本 |
| c | 字符串 | 是 | 无 | 客户端标识符 |
| f | 字符串 | 否 | xml | 响应格式 |

**请求示例：**

```
http://server/rest/startScan.view?u=admin&t=token&s=salt&v=1.16.1&c=client
```

---

## 4. 错误码说明

### 错误结构

```json
{
  "subsonic-response": {
    "status": "failed",
    "version": "1.16.1",
    "error": {
      "code": 40, 
      "message": "Unauthorized"
    }
  }
}
```

### 常见错误码

| 错误码 | 描述 |
|--------|------|
| 0 | 一般错误 |
| 10 | 所需参数缺失 |
| 20 | 不支持的操作 |
| 30 | 未找到 |
| 40 | 未授权 |
| 50 | 服务器错误 |
| 60 | 更新数据库失败 |
| 70 | 参数错误 |

---

## 5. 扩展能力

OpenSubsonic API 提供了一些扩展能力，通过 `getOpenSubsonicExtensions` 方法可以列出服务器支持的 OpenSubsonic 扩展。

这些扩展可能包括：
- 同步歌词支持
- 多语言歌词
- 按歌曲 ID 检索歌词
- 其他自定义功能

具体支持的扩展取决于服务器实现。