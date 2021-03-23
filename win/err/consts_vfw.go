package err

const (
	VFW_S_NO_MORE_ITEMS                      ERROR = 0x00040103
	VFW_S_DUPLICATE_NAME                     ERROR = 0x0004022d
	VFW_S_STATE_INTERMEDIATE                 ERROR = 0x00040237
	VFW_S_PARTIAL_RENDER                     ERROR = 0x00040242
	VFW_S_SOME_DATA_IGNORED                  ERROR = 0x00040245
	VFW_S_CONNECTIONS_DEFERRED               ERROR = 0x00040246
	VFW_S_RESOURCE_NOT_NEEDED                ERROR = 0x00040250
	VFW_S_MEDIA_TYPE_IGNORED                 ERROR = 0x00040254
	VFW_S_VIDEO_NOT_RENDERED                 ERROR = 0x00040257
	VFW_S_AUDIO_NOT_RENDERED                 ERROR = 0x00040258
	VFW_S_RPZA                               ERROR = 0x0004025a
	VFW_S_ESTIMATED                          ERROR = 0x00040260
	VFW_S_RESERVED                           ERROR = 0x00040263
	VFW_S_STREAM_OFF                         ERROR = 0x00040267
	VFW_S_CANT_CUE                           ERROR = 0x00040268
	VFW_S_NOPREVIEWPIN                       ERROR = 0x0004027e
	VFW_S_DVD_NON_ONE_SEQUENTIAL             ERROR = 0x00040280
	VFW_S_DVD_CHANNEL_CONTENTS_NOT_AVAILABLE ERROR = 0x0004028c
	VFW_S_DVD_NOT_ACCURATE                   ERROR = 0x0004028d
	VFW_E_INVALIDMEDIATYPE                   ERROR = 0x80040200
	VFW_E_INVALIDSUBTYPE                     ERROR = 0x80040201
	VFW_E_NEED_OWNER                         ERROR = 0x80040202
	VFW_E_ENUM_OUT_OF_SYNC                   ERROR = 0x80040203
	VFW_E_ALREADY_CONNECTED                  ERROR = 0x80040204
	VFW_E_FILTER_ACTIVE                      ERROR = 0x80040205
	VFW_E_NO_TYPES                           ERROR = 0x80040206
	VFW_E_NO_ACCEPTABLE_TYPES                ERROR = 0x80040207
	VFW_E_INVALID_DIRECTION                  ERROR = 0x80040208
	VFW_E_NOT_CONNECTED                      ERROR = 0x80040209
	VFW_E_NO_ALLOCATOR                       ERROR = 0x8004020a
	VFW_E_RUNTIME_ERROR                      ERROR = 0x8004020b
	VFW_E_BUFFER_NOTSET                      ERROR = 0x8004020c
	VFW_E_BUFFER_OVERFLOW                    ERROR = 0x8004020d
	VFW_E_BADALIGN                           ERROR = 0x8004020e
	VFW_E_ALREADY_COMMITTED                  ERROR = 0x8004020f
	VFW_E_BUFFERS_OUTSTANDING                ERROR = 0x80040210
	VFW_E_NOT_COMMITTED                      ERROR = 0x80040211
	VFW_E_SIZENOTSET                         ERROR = 0x80040212
	VFW_E_NO_CLOCK                           ERROR = 0x80040213
	VFW_E_NO_SINK                            ERROR = 0x80040214
	VFW_E_NO_INTERFACE                       ERROR = 0x80040215
	VFW_E_NOT_FOUND                          ERROR = 0x80040216
	VFW_E_CANNOT_CONNECT                     ERROR = 0x80040217
	VFW_E_CANNOT_RENDER                      ERROR = 0x80040218
	VFW_E_CHANGING_FORMAT                    ERROR = 0x80040219
	VFW_E_NO_COLOR_KEY_SET                   ERROR = 0x8004021a
	VFW_E_NOT_OVERLAY_CONNECTION             ERROR = 0x8004021b
	VFW_E_NOT_SAMPLE_CONNECTION              ERROR = 0x8004021c
	VFW_E_PALETTE_SET                        ERROR = 0x8004021d
	VFW_E_COLOR_KEY_SET                      ERROR = 0x8004021e
	VFW_E_NO_COLOR_KEY_FOUND                 ERROR = 0x8004021f
	VFW_E_NO_PALETTE_AVAILABLE               ERROR = 0x80040220
	VFW_E_NO_DISPLAY_PALETTE                 ERROR = 0x80040221
	VFW_E_TOO_MANY_COLORS                    ERROR = 0x80040222
	VFW_E_STATE_CHANGED                      ERROR = 0x80040223
	VFW_E_NOT_STOPPED                        ERROR = 0x80040224
	VFW_E_NOT_PAUSED                         ERROR = 0x80040225
	VFW_E_NOT_RUNNING                        ERROR = 0x80040226
	VFW_E_WRONG_STATE                        ERROR = 0x80040227
	VFW_E_START_TIME_AFTER_END               ERROR = 0x80040228
	VFW_E_INVALID_RECT                       ERROR = 0x80040229
	VFW_E_TYPE_NOT_ACCEPTED                  ERROR = 0x8004022a
	VFW_E_SAMPLE_REJECTED                    ERROR = 0x8004022b
	VFW_E_SAMPLE_REJECTED_EOS                ERROR = 0x8004022c
	VFW_E_DUPLICATE_NAME                     ERROR = 0x8004022d
	VFW_E_TIMEOUT                            ERROR = 0x8004022e
	VFW_E_INVALID_FILE_FORMAT                ERROR = 0x8004022f
	VFW_E_ENUM_OUT_OF_RANGE                  ERROR = 0x80040230
	VFW_E_CIRCULAR_GRAPH                     ERROR = 0x80040231
	VFW_E_NOT_ALLOWED_TO_SAVE                ERROR = 0x80040232
	VFW_E_TIME_ALREADY_PASSED                ERROR = 0x80040233
	VFW_E_ALREADY_CANCELLED                  ERROR = 0x80040234
	VFW_E_CORRUPT_GRAPH_FILE                 ERROR = 0x80040235
	VFW_E_ADVISE_ALREADY_SET                 ERROR = 0x80040236
	VFW_E_NO_MODEX_AVAILABLE                 ERROR = 0x80040238
	VFW_E_NO_ADVISE_SET                      ERROR = 0x80040239
	VFW_E_NO_FULLSCREEN                      ERROR = 0x8004023a
	VFW_E_IN_FULLSCREEN_MODE                 ERROR = 0x8004023b
	VFW_E_UNKNOWN_FILE_TYPE                  ERROR = 0x80040240
	VFW_E_CANNOT_LOAD_SOURCE_FILTER          ERROR = 0x80040241
	VFW_E_FILE_TOO_SHORT                     ERROR = 0x80040243
	VFW_E_INVALID_FILE_VERSION               ERROR = 0x80040244
	VFW_E_INVALID_CLSID                      ERROR = 0x80040247
	VFW_E_INVALID_MEDIA_TYPE                 ERROR = 0x80040248
	VFW_E_SAMPLE_TIME_NOT_SET                ERROR = 0x80040249
	VFW_E_MEDIA_TIME_NOT_SET                 ERROR = 0x80040251
	VFW_E_NO_TIME_FORMAT_SET                 ERROR = 0x80040252
	VFW_E_MONO_AUDIO_HW                      ERROR = 0x80040253
	VFW_E_NO_DECOMPRESSOR                    ERROR = 0x80040255
	VFW_E_NO_AUDIO_HARDWARE                  ERROR = 0x80040256
	VFW_E_RPZA                               ERROR = 0x80040259
	VFW_E_PROCESSOR_NOT_SUITABLE             ERROR = 0x8004025b
	VFW_E_UNSUPPORTED_AUDIO                  ERROR = 0x8004025c
	VFW_E_UNSUPPORTED_VIDEO                  ERROR = 0x8004025d
	VFW_E_MPEG_NOT_CONSTRAINED               ERROR = 0x8004025e
	VFW_E_NOT_IN_GRAPH                       ERROR = 0x8004025f
	VFW_E_NO_TIME_FORMAT                     ERROR = 0x80040261
	VFW_E_READ_ONLY                          ERROR = 0x80040262
	VFW_E_BUFFER_UNDERFLOW                   ERROR = 0x80040264
	VFW_E_UNSUPPORTED_STREAM                 ERROR = 0x80040265
	VFW_E_NO_TRANSPORT                       ERROR = 0x80040266
	VFW_E_BAD_VIDEOCD                        ERROR = 0x80040269
	VFW_S_NO_STOP_TIME                       ERROR = 0x80040270
	VFW_E_OUT_OF_VIDEO_MEMORY                ERROR = 0x80040271
	VFW_E_VP_NEGOTIATION_FAILED              ERROR = 0x80040272
	VFW_E_DDRAW_CAPS_NOT_SUITABLE            ERROR = 0x80040273
	VFW_E_NO_VP_HARDWARE                     ERROR = 0x80040274
	VFW_E_NO_CAPTURE_HARDWARE                ERROR = 0x80040275
	VFW_E_DVD_OPERATION_INHIBITED            ERROR = 0x80040276
	VFW_E_DVD_INVALIDDOMAIN                  ERROR = 0x80040277
	VFW_E_DVD_NO_BUTTON                      ERROR = 0x80040278
	VFW_E_DVD_GRAPHNOTREADY                  ERROR = 0x80040279
	VFW_E_DVD_RENDERFAIL                     ERROR = 0x8004027a
	VFW_E_DVD_DECNOTENOUGH                   ERROR = 0x8004027b
	VFW_E_DDRAW_VERSION_NOT_SUITABLE         ERROR = 0x8004027c
	VFW_E_COPYPROT_FAILED                    ERROR = 0x8004027d
	VFW_E_TIME_EXPIRED                       ERROR = 0x8004027f
	VFW_E_DVD_WRONG_SPEED                    ERROR = 0x80040281
	VFW_E_DVD_MENU_DOES_NOT_EXIST            ERROR = 0x80040282
	VFW_E_DVD_CMD_CANCELLED                  ERROR = 0x80040283
	VFW_E_DVD_STATE_WRONG_VERSION            ERROR = 0x80040284
	VFW_E_DVD_STATE_CORRUPT                  ERROR = 0x80040285
	VFW_E_DVD_STATE_WRONG_DISC               ERROR = 0x80040286
	VFW_E_DVD_INCOMPATIBLE_REGION            ERROR = 0x80040287
	VFW_E_DVD_NO_ATTRIBUTES                  ERROR = 0x80040288
	VFW_E_DVD_NO_GOUP_PGC                    ERROR = 0x80040289
	VFW_E_DVD_LOW_PARENTAL_LEVEL             ERROR = 0x8004028a
	VFW_E_DVD_NOT_IN_KARAOKE_MODE            ERROR = 0x8004028b
	VFW_E_FRAME_STEP_UNSUPPORTED             ERROR = 0x8004028e
	VFW_E_DVD_STREAM_DISABLED                ERROR = 0x8004028f
	VFW_E_DVD_TITLE_UNKNOWN                  ERROR = 0x80040290
	VFW_E_DVD_INVALID_DISC                   ERROR = 0x80040291
	VFW_E_DVD_NO_RESUME_INFORMATION          ERROR = 0x80040292
	VFW_E_PIN_ALREADY_BLOCKED_ON_THIS_THREAD ERROR = 0x80040293
	VFW_E_PIN_ALREADY_BLOCKED                ERROR = 0x80040294
	VFW_E_CERTIFICATION_FAILURE              ERROR = 0x80040295
	VFW_E_VMR_NOT_IN_MIXER_MODE              ERROR = 0x80040296
	VFW_E_VMR_NO_AP_SUPPLIED                 ERROR = 0x80040297
	VFW_E_VMR_NO_DEINTERLACE_HW              ERROR = 0x80040298
	VFW_E_VMR_NO_PROCAMP_HW                  ERROR = 0x80040299
	VFW_E_DVD_VMR9_INCOMPATIBLEDEC           ERROR = 0x8004029a
	VFW_E_NO_COPP_HW                         ERROR = 0x8004029b
	VFW_E_BAD_KEY                            ERROR = 0x800403f2
)
