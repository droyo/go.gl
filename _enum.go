package gl

// GoGL2 - automatically generated OpenGL binding: http://github.com/chsc/gogl2

const (
	gl_2D = 0x0600
	gl_2_BYTES = 0x1407
	gl_3D = 0x0601
	gl_3D_COLOR = 0x0602
	gl_3D_COLOR_TEXTURE = 0x0603
	gl_3_BYTES = 0x1408
	gl_4D_COLOR_TEXTURE = 0x0604
	gl_4_BYTES = 0x1409
	gl_ACCUM = 0x0100
	gl_ACCUM_ALPHA_BITS = 0x0D5B
	gl_ACCUM_BLUE_BITS = 0x0D5A
	gl_ACCUM_BUFFER_BIT = 0x00000200
	gl_ACCUM_CLEAR_VALUE = 0x0B80
	gl_ACCUM_GREEN_BITS = 0x0D59
	gl_ACCUM_RED_BITS = 0x0D58
	gl_ACTIVE_ATTRIBUTES = 0x8B89
	gl_ACTIVE_ATTRIBUTE_MAX_LENGTH = 0x8B8A
	gl_ACTIVE_TEXTURE = 0x84E0
	gl_ACTIVE_UNIFORMS = 0x8B86
	gl_ACTIVE_UNIFORM_MAX_LENGTH = 0x8B87
	gl_ADD = 0x0104
	gl_ADD_SIGNED = 0x8574
	gl_ALIASED_LINE_WIDTH_RANGE = 0x846E
	gl_ALIASED_POINT_SIZE_RANGE = 0x846D
	gl_ALL_ATTRIB_BITS = 0xFFFFFFFF
	gl_ALPHA = 0x1906
	gl_ALPHA12 = 0x803D
	gl_ALPHA16 = 0x803E
	gl_ALPHA4 = 0x803B
	gl_ALPHA8 = 0x803C
	gl_ALPHA_BIAS = 0x0D1D
	gl_ALPHA_BITS = 0x0D55
	gl_ALPHA_SCALE = 0x0D1C
	gl_ALPHA_TEST = 0x0BC0
	gl_ALPHA_TEST_FUNC = 0x0BC1
	gl_ALPHA_TEST_REF = 0x0BC2
	gl_ALWAYS = 0x0207
	gl_AMBIENT = 0x1200
	gl_AMBIENT_AND_DIFFUSE = 0x1602
	gl_AND = 0x1501
	gl_AND_INVERTED = 0x1504
	gl_AND_REVERSE = 0x1502
	gl_ARRAY_BUFFER = 0x8892
	gl_ARRAY_BUFFER_BINDING = 0x8894
	gl_ATTACHED_SHADERS = 0x8B85
	gl_ATTRIB_STACK_DEPTH = 0x0BB0
	gl_AUTO_NORMAL = 0x0D80
	gl_AUX0 = 0x0409
	gl_AUX1 = 0x040A
	gl_AUX2 = 0x040B
	gl_AUX3 = 0x040C
	gl_AUX_BUFFERS = 0x0C00
	gl_BACK = 0x0405
	gl_BACK_LEFT = 0x0402
	gl_BACK_RIGHT = 0x0403
	gl_BGR = 0x80E0
	gl_BGRA = 0x80E1
	gl_BITMAP = 0x1A00
	gl_BITMAP_TOKEN = 0x0704
	gl_BLEND_COLOR = 0x8005
	gl_BLEND_DST = 0x0BE0
	gl_BLEND_DST_ALPHA = 0x80CA
	gl_BLEND_DST_RGB = 0x80C8
	gl_BLEND_EQUATION = 0x8009
	gl_BLEND_EQUATION_ALPHA = 0x883D
	gl_BLEND_EQUATION_RGB = 0x8009
	gl_BLEND_SRC = 0x0BE1
	gl_BLEND_SRC_ALPHA = 0x80CB
	gl_BLEND_SRC_RGB = 0x80C9
	gl_BLUE = 0x1905
	gl_BLUE_BIAS = 0x0D1B
	gl_BLUE_BITS = 0x0D54
	gl_BLUE_SCALE = 0x0D1A
	gl_BOOL = 0x8B56
	gl_BOOL_VEC2 = 0x8B57
	gl_BOOL_VEC3 = 0x8B58
	gl_BOOL_VEC4 = 0x8B59
	gl_BUFFER_ACCESS = 0x88BB
	gl_BUFFER_MAPPED = 0x88BC
	gl_BUFFER_MAP_POINTER = 0x88BD
	gl_BUFFER_SIZE = 0x8764
	gl_BUFFER_USAGE = 0x8765
	gl_BYTE = 0x1400
	gl_C3F_V3F = 0x2A24
	gl_C4F_N3F_V3F = 0x2A26
	gl_C4UB_V2F = 0x2A22
	gl_C4UB_V3F = 0x2A23
	gl_CCW = 0x0901
	gl_CLAMP = 0x2900
	gl_CLAMP_TO_BORDER = 0x812D
	gl_CLAMP_TO_EDGE = 0x812F
	gl_CLEAR = 0x1500
	gl_CLIENT_ACTIVE_TEXTURE = 0x84E1
	gl_CLIENT_ALL_ATTRIB_BITS = 0xFFFFFFFF
	gl_CLIENT_ATTRIB_STACK_DEPTH = 0x0BB1
	gl_CLIENT_PIXEL_STORE_BIT = 0x00000001
	gl_CLIENT_VERTEX_ARRAY_BIT = 0x00000002
	gl_CLIP_PLANE0 = 0x3000
	gl_CLIP_PLANE1 = 0x3001
	gl_CLIP_PLANE2 = 0x3002
	gl_CLIP_PLANE3 = 0x3003
	gl_CLIP_PLANE4 = 0x3004
	gl_CLIP_PLANE5 = 0x3005
	gl_COEFF = 0x0A00
	gl_COLOR = 0x1800
	gl_COLOR_ARRAY = 0x8076
	gl_COLOR_ARRAY_BUFFER_BINDING = 0x8898
	gl_COLOR_ARRAY_POINTER = 0x8090
	gl_COLOR_ARRAY_SIZE = 0x8081
	gl_COLOR_ARRAY_STRIDE = 0x8083
	gl_COLOR_ARRAY_TYPE = 0x8082
	gl_COLOR_ATTACHMENT0 = 0x8CE0
	gl_COLOR_BUFFER_BIT = 0x00004000
	gl_COLOR_CLEAR_VALUE = 0x0C22
	gl_COLOR_INDEX = 0x1900
	gl_COLOR_INDEXES = 0x1603
	gl_COLOR_MATERIAL = 0x0B57
	gl_COLOR_MATERIAL_FACE = 0x0B55
	gl_COLOR_MATERIAL_PARAMETER = 0x0B56
	gl_COLOR_SUM = 0x8458
	gl_COLOR_WRITEMASK = 0x0C23
	gl_COMBINE = 0x8570
	gl_COMBINE_ALPHA = 0x8572
	gl_COMBINE_RGB = 0x8571
	gl_COMPARE_R_TO_TEXTURE = 0x884E
	gl_COMPILE = 0x1300
	gl_COMPILE_AND_EXECUTE = 0x1301
	gl_COMPILE_STATUS = 0x8B81
	gl_COMPRESSED_ALPHA = 0x84E9
	gl_COMPRESSED_INTENSITY = 0x84EC
	gl_COMPRESSED_LUMINANCE = 0x84EA
	gl_COMPRESSED_LUMINANCE_ALPHA = 0x84EB
	gl_COMPRESSED_RGB = 0x84ED
	gl_COMPRESSED_RGBA = 0x84EE
	gl_COMPRESSED_TEXTURE_FORMATS = 0x86A3
	gl_CONSTANT = 0x8576
	gl_CONSTANT_ALPHA = 0x8003
	gl_CONSTANT_ATTENUATION = 0x1207
	gl_CONSTANT_COLOR = 0x8001
	gl_COORD_REPLACE = 0x8862
	gl_COPY = 0x1503
	gl_COPY_INVERTED = 0x150C
	gl_COPY_PIXEL_TOKEN = 0x0706
	gl_CULL_FACE_MODE = 0x0B45
	gl_CURRENT_BIT = 0x00000001
	gl_CURRENT_COLOR = 0x0B00
	gl_CURRENT_FOG_COORD = 0x8453
	gl_CURRENT_FOG_COORDINATE = 0x8453
	gl_CURRENT_INDEX = 0x0B01
	gl_CURRENT_NORMAL = 0x0B02
	gl_CURRENT_PROGRAM = 0x8B8D
	gl_CURRENT_QUERY = 0x8865
	gl_CURRENT_RASTER_COLOR = 0x0B04
	gl_CURRENT_RASTER_DISTANCE = 0x0B09
	gl_CURRENT_RASTER_INDEX = 0x0B05
	gl_CURRENT_RASTER_POSITION = 0x0B07
	gl_CURRENT_RASTER_POSITION_VALID = 0x0B08
	gl_CURRENT_RASTER_TEXTURE_COORDS = 0x0B06
	gl_CURRENT_SECONDARY_COLOR = 0x8459
	gl_CURRENT_TEXTURE_COORDS = 0x0B03
	gl_CURRENT_VERTEX_ATTRIB = 0x8626
	gl_CW = 0x0900
	gl_DECAL = 0x2101
	gl_DECR = 0x1E03
	gl_DECR_WRAP = 0x8508
	gl_DELETE_STATUS = 0x8B80
	gl_DEPTH = 0x1801
	gl_DEPTH_ATTACHMENT = 0x8D00
	gl_DEPTH_BIAS = 0x0D1F
	gl_DEPTH_BITS = 0x0D56
	gl_DEPTH_BUFFER_BIT = 0x00000100
	gl_DEPTH_CLEAR_VALUE = 0x0B73
	gl_DEPTH_COMPONENT = 0x1902
	gl_DEPTH_COMPONENT16 = 0x81A5
	gl_DEPTH_COMPONENT24 = 0x81A6
	gl_DEPTH_COMPONENT32 = 0x81A7
	gl_DEPTH_FUNC = 0x0B74
	gl_DEPTH_RANGE = 0x0B70
	gl_DEPTH_SCALE = 0x0D1E
	gl_DEPTH_TEXTURE_MODE = 0x884B
	gl_DEPTH_WRITEMASK = 0x0B72
	gl_DIFFUSE = 0x1201
	gl_DOMAIN = 0x0A02
	gl_DONT_CARE = 0x1100
	gl_DOT3_RGB = 0x86AE
	gl_DOT3_RGBA = 0x86AF
	gl_DOUBLE = 0x140A
	gl_DOUBLEBUFFER = 0x0C32
	gl_DRAW_BUFFER = 0x0C01
	gl_DRAW_BUFFER0 = 0x8825
	gl_DRAW_BUFFER1 = 0x8826
	gl_DRAW_BUFFER10 = 0x882F
	gl_DRAW_BUFFER11 = 0x8830
	gl_DRAW_BUFFER12 = 0x8831
	gl_DRAW_BUFFER13 = 0x8832
	gl_DRAW_BUFFER14 = 0x8833
	gl_DRAW_BUFFER15 = 0x8834
	gl_DRAW_BUFFER2 = 0x8827
	gl_DRAW_BUFFER3 = 0x8828
	gl_DRAW_BUFFER4 = 0x8829
	gl_DRAW_BUFFER5 = 0x882A
	gl_DRAW_BUFFER6 = 0x882B
	gl_DRAW_BUFFER7 = 0x882C
	gl_DRAW_BUFFER8 = 0x882D
	gl_DRAW_BUFFER9 = 0x882E
	gl_DRAW_PIXEL_TOKEN = 0x0705
	gl_DST_ALPHA = 0x0304
	gl_DST_COLOR = 0x0306
	gl_DYNAMIC_COPY = 0x88EA
	gl_DYNAMIC_DRAW = 0x88E8
	gl_DYNAMIC_READ = 0x88E9
	gl_EDGE_FLAG = 0x0B43
	gl_EDGE_FLAG_ARRAY = 0x8079
	gl_EDGE_FLAG_ARRAY_BUFFER_BINDING = 0x889B
	gl_EDGE_FLAG_ARRAY_POINTER = 0x8093
	gl_EDGE_FLAG_ARRAY_STRIDE = 0x808C
	gl_ELEMENT_ARRAY_BUFFER = 0x8893
	gl_ELEMENT_ARRAY_BUFFER_BINDING = 0x8895
	gl_EMISSION = 0x1600
	gl_ENABLE_BIT = 0x00002000
	gl_EQUAL = 0x0202
	gl_EQUIV = 0x1509
	gl_EVAL_BIT = 0x00010000
	gl_EXP = 0x0800
	gl_EXP2 = 0x0801
	gl_EXTENSIONS = 0x1F03
	gl_EYE_LINEAR = 0x2400
	gl_EYE_PLANE = 0x2502
	gl_FALSE = 0
	gl_FASTEST = 0x1101
	gl_FEEDBACK = 0x1C01
	gl_FEEDBACK_BUFFER_POINTER = 0x0DF0
	gl_FEEDBACK_BUFFER_SIZE = 0x0DF1
	gl_FEEDBACK_BUFFER_TYPE = 0x0DF2
	gl_FILL = 0x1B02
	gl_FIXED = 0x140C
	gl_FLAT = 0x1D00
	gl_FLOAT = 0x1406
	gl_FLOAT_MAT2 = 0x8B5A
	gl_FLOAT_MAT3 = 0x8B5B
	gl_FLOAT_MAT4 = 0x8B5C
	gl_FLOAT_VEC2 = 0x8B50
	gl_FLOAT_VEC3 = 0x8B51
	gl_FLOAT_VEC4 = 0x8B52
	gl_FOG = 0x0B60
	gl_FOG_BIT = 0x00000080
	gl_FOG_COLOR = 0x0B66
	gl_FOG_COORD = 0x8451
	gl_FOG_COORDINATE = 0x8451
	gl_FOG_COORDINATE_ARRAY = 0x8457
	gl_FOG_COORDINATE_ARRAY_BUFFER_BINDING = 0x889D
	gl_FOG_COORDINATE_ARRAY_POINTER = 0x8456
	gl_FOG_COORDINATE_ARRAY_STRIDE = 0x8455
	gl_FOG_COORDINATE_ARRAY_TYPE = 0x8454
	gl_FOG_COORDINATE_SOURCE = 0x8450
	gl_FOG_COORD_ARRAY = 0x8457
	gl_FOG_COORD_ARRAY_BUFFER_BINDING = 0x889D
	gl_FOG_COORD_ARRAY_POINTER = 0x8456
	gl_FOG_COORD_ARRAY_STRIDE = 0x8455
	gl_FOG_COORD_ARRAY_TYPE = 0x8454
	gl_FOG_COORD_SRC = 0x8450
	gl_FOG_DENSITY = 0x0B62
	gl_FOG_END = 0x0B64
	gl_FOG_HINT = 0x0C54
	gl_FOG_INDEX = 0x0B61
	gl_FOG_MODE = 0x0B65
	gl_FOG_START = 0x0B63
	gl_FRAGMENT_DEPTH = 0x8452
	gl_FRAGMENT_SHADER = 0x8B30
	gl_FRAGMENT_SHADER_DERIVATIVE_HINT = 0x8B8B
	gl_FRAMEBUFFER = 0x8D40
	gl_FRAMEBUFFER_ATTACHMENT_OBJECT_NAME = 0x8CD1
	gl_FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE = 0x8CD0
	gl_FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE = 0x8CD3
	gl_FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL = 0x8CD2
	gl_FRAMEBUFFER_BINDING = 0x8CA6
	gl_FRAMEBUFFER_COMPLETE = 0x8CD5
	gl_FRAMEBUFFER_INCOMPLETE_ATTACHMENT = 0x8CD6
	gl_FRAMEBUFFER_INCOMPLETE_DIMENSIONS = 0x8CD9
	gl_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT = 0x8CD7
	gl_FRAMEBUFFER_UNSUPPORTED = 0x8CDD
	gl_FRONT = 0x0404
	gl_FRONT_AND_BACK = 0x0408
	gl_FRONT_FACE = 0x0B46
	gl_FRONT_LEFT = 0x0400
	gl_FRONT_RIGHT = 0x0401
	gl_FUNC_ADD = 0x8006
	gl_FUNC_REVERSE_SUBTRACT = 0x800B
	gl_FUNC_SUBTRACT = 0x800A
	gl_GENERATE_MIPMAP = 0x8191
	gl_GENERATE_MIPMAP_HINT = 0x8192
	gl_GEQUAL = 0x0206
	gl_GREATER = 0x0204
	gl_GREEN = 0x1904
	gl_GREEN_BIAS = 0x0D19
	gl_GREEN_BITS = 0x0D53
	gl_GREEN_SCALE = 0x0D18
	gl_HIGH_FLOAT = 0x8DF2
	gl_HIGH_INT = 0x8DF5
	gl_HINT_BIT = 0x00008000
	gl_IMPLEMENTATION_COLOR_READ_FORMAT = 0x8B9B
	gl_IMPLEMENTATION_COLOR_READ_TYPE = 0x8B9A
	gl_INCR = 0x1E02
	gl_INCR_WRAP = 0x8507
	gl_INDEX_ARRAY = 0x8077
	gl_INDEX_ARRAY_BUFFER_BINDING = 0x8899
	gl_INDEX_ARRAY_POINTER = 0x8091
	gl_INDEX_ARRAY_STRIDE = 0x8086
	gl_INDEX_ARRAY_TYPE = 0x8085
	gl_INDEX_BITS = 0x0D51
	gl_INDEX_CLEAR_VALUE = 0x0C20
	gl_INDEX_LOGIC_OP = 0x0BF1
	gl_INDEX_MODE = 0x0C30
	gl_INDEX_OFFSET = 0x0D13
	gl_INDEX_SHIFT = 0x0D12
	gl_INDEX_WRITEMASK = 0x0C21
	gl_INFO_LOG_LENGTH = 0x8B84
	gl_INT = 0x1404
	gl_INTENSITY = 0x8049
	gl_INTENSITY12 = 0x804C
	gl_INTENSITY16 = 0x804D
	gl_INTENSITY4 = 0x804A
	gl_INTENSITY8 = 0x804B
	gl_INTERPOLATE = 0x8575
	gl_INT_VEC2 = 0x8B53
	gl_INT_VEC3 = 0x8B54
	gl_INT_VEC4 = 0x8B55
	gl_INVALID_ENUM = 0x0500
	gl_INVALID_FRAMEBUFFER_OPERATION = 0x0506
	gl_INVALID_OPERATION = 0x0502
	gl_INVALID_VALUE = 0x0501
	gl_INVERT = 0x150A
	gl_KEEP = 0x1E00
	gl_LEFT = 0x0406
	gl_LEQUAL = 0x0203
	gl_LESS = 0x0201
	gl_LIGHT0 = 0x4000
	gl_LIGHT1 = 0x4001
	gl_LIGHT2 = 0x4002
	gl_LIGHT3 = 0x4003
	gl_LIGHT4 = 0x4004
	gl_LIGHT5 = 0x4005
	gl_LIGHT6 = 0x4006
	gl_LIGHT7 = 0x4007
	gl_LIGHTING = 0x0B50
	gl_LIGHTING_BIT = 0x00000040
	gl_LIGHT_MODEL_AMBIENT = 0x0B53
	gl_LIGHT_MODEL_COLOR_CONTROL = 0x81F8
	gl_LIGHT_MODEL_LOCAL_VIEWER = 0x0B51
	gl_LIGHT_MODEL_TWO_SIDE = 0x0B52
	gl_LINE = 0x1B01
	gl_LINEAR = 0x2601
	gl_LINEAR_ATTENUATION = 0x1208
	gl_LINEAR_MIPMAP_LINEAR = 0x2703
	gl_LINEAR_MIPMAP_NEAREST = 0x2701
	gl_LINES = 0x0001
	gl_LINE_BIT = 0x00000004
	gl_LINE_LOOP = 0x0002
	gl_LINE_RESET_TOKEN = 0x0707
	gl_LINE_SMOOTH = 0x0B20
	gl_LINE_SMOOTH_HINT = 0x0C52
	gl_LINE_STIPPLE = 0x0B24
	gl_LINE_STIPPLE_PATTERN = 0x0B25
	gl_LINE_STIPPLE_REPEAT = 0x0B26
	gl_LINE_STRIP = 0x0003
	gl_LINE_TOKEN = 0x0702
	gl_LINE_WIDTH = 0x0B21
	gl_LINE_WIDTH_GRANULARITY = 0x0B23
	gl_LINE_WIDTH_RANGE = 0x0B22
	gl_LINK_STATUS = 0x8B82
	gl_LIST_BASE = 0x0B32
	gl_LIST_BIT = 0x00020000
	gl_LIST_INDEX = 0x0B33
	gl_LIST_MODE = 0x0B30
	gl_LOAD = 0x0101
	gl_LOGIC_OP = 0x0BF1
	gl_LOGIC_OP_MODE = 0x0BF0
	gl_LOWER_LEFT = 0x8CA1
	gl_LOW_FLOAT = 0x8DF0
	gl_LOW_INT = 0x8DF3
	gl_LUMINANCE = 0x1909
	gl_LUMINANCE12 = 0x8041
	gl_LUMINANCE12_ALPHA12 = 0x8047
	gl_LUMINANCE12_ALPHA4 = 0x8046
	gl_LUMINANCE16 = 0x8042
	gl_LUMINANCE16_ALPHA16 = 0x8048
	gl_LUMINANCE4 = 0x803F
	gl_LUMINANCE4_ALPHA4 = 0x8043
	gl_LUMINANCE6_ALPHA2 = 0x8044
	gl_LUMINANCE8 = 0x8040
	gl_LUMINANCE8_ALPHA8 = 0x8045
	gl_LUMINANCE_ALPHA = 0x190A
	gl_MAP1_COLOR_4 = 0x0D90
	gl_MAP1_GRID_DOMAIN = 0x0DD0
	gl_MAP1_GRID_SEGMENTS = 0x0DD1
	gl_MAP1_INDEX = 0x0D91
	gl_MAP1_NORMAL = 0x0D92
	gl_MAP1_TEXTURE_COORD_1 = 0x0D93
	gl_MAP1_TEXTURE_COORD_2 = 0x0D94
	gl_MAP1_TEXTURE_COORD_3 = 0x0D95
	gl_MAP1_TEXTURE_COORD_4 = 0x0D96
	gl_MAP1_VERTEX_3 = 0x0D97
	gl_MAP1_VERTEX_4 = 0x0D98
	gl_MAP2_COLOR_4 = 0x0DB0
	gl_MAP2_GRID_DOMAIN = 0x0DD2
	gl_MAP2_GRID_SEGMENTS = 0x0DD3
	gl_MAP2_INDEX = 0x0DB1
	gl_MAP2_NORMAL = 0x0DB2
	gl_MAP2_TEXTURE_COORD_1 = 0x0DB3
	gl_MAP2_TEXTURE_COORD_2 = 0x0DB4
	gl_MAP2_TEXTURE_COORD_3 = 0x0DB5
	gl_MAP2_TEXTURE_COORD_4 = 0x0DB6
	gl_MAP2_VERTEX_3 = 0x0DB7
	gl_MAP2_VERTEX_4 = 0x0DB8
	gl_MAP_COLOR = 0x0D10
	gl_MAP_STENCIL = 0x0D11
	gl_MATRIX_MODE = 0x0BA0
	gl_MAX = 0x8008
	gl_MAX_3D_TEXTURE_SIZE = 0x8073
	gl_MAX_ATTRIB_STACK_DEPTH = 0x0D35
	gl_MAX_CLIENT_ATTRIB_STACK_DEPTH = 0x0D3B
	gl_MAX_CLIP_PLANES = 0x0D32
	gl_MAX_COMBINED_TEXTURE_IMAGE_UNITS = 0x8B4D
	gl_MAX_CUBE_MAP_TEXTURE_SIZE = 0x851C
	gl_MAX_DRAW_BUFFERS = 0x8824
	gl_MAX_ELEMENTS_INDICES = 0x80E9
	gl_MAX_ELEMENTS_VERTICES = 0x80E8
	gl_MAX_EVAL_ORDER = 0x0D30
	gl_MAX_FRAGMENT_UNIFORM_COMPONENTS = 0x8B49
	gl_MAX_FRAGMENT_UNIFORM_VECTORS = 0x8DFD
	gl_MAX_LIGHTS = 0x0D31
	gl_MAX_LIST_NESTING = 0x0B31
	gl_MAX_MODELVIEW_STACK_DEPTH = 0x0D36
	gl_MAX_NAME_STACK_DEPTH = 0x0D37
	gl_MAX_PIXEL_MAP_TABLE = 0x0D34
	gl_MAX_PROJECTION_STACK_DEPTH = 0x0D38
	gl_MAX_RENDERBUFFER_SIZE = 0x84E8
	gl_MAX_TEXTURE_COORDS = 0x8871
	gl_MAX_TEXTURE_IMAGE_UNITS = 0x8872
	gl_MAX_TEXTURE_LOD_BIAS = 0x84FD
	gl_MAX_TEXTURE_SIZE = 0x0D33
	gl_MAX_TEXTURE_STACK_DEPTH = 0x0D39
	gl_MAX_TEXTURE_UNITS = 0x84E2
	gl_MAX_VARYING_FLOATS = 0x8B4B
	gl_MAX_VARYING_VECTORS = 0x8DFC
	gl_MAX_VERTEX_ATTRIBS = 0x8869
	gl_MAX_VERTEX_TEXTURE_IMAGE_UNITS = 0x8B4C
	gl_MAX_VERTEX_UNIFORM_COMPONENTS = 0x8B4A
	gl_MAX_VERTEX_UNIFORM_VECTORS = 0x8DFB
	gl_MAX_VIEWPORT_DIMS = 0x0D3A
	gl_MEDIUM_FLOAT = 0x8DF1
	gl_MEDIUM_INT = 0x8DF4
	gl_MIN = 0x8007
	gl_MIRRORED_REPEAT = 0x8370
	gl_MODELVIEW = 0x1700
	gl_MODELVIEW_MATRIX = 0x0BA6
	gl_MODELVIEW_STACK_DEPTH = 0x0BA3
	gl_MODULATE = 0x2100
	gl_MULT = 0x0103
	gl_MULTISAMPLE = 0x809D
	gl_MULTISAMPLE_BIT = 0x20000000
	gl_N3F_V3F = 0x2A25
	gl_NAME_STACK_DEPTH = 0x0D70
	gl_NAND = 0x150E
	gl_NEAREST = 0x2600
	gl_NEAREST_MIPMAP_LINEAR = 0x2702
	gl_NEAREST_MIPMAP_NEAREST = 0x2700
	gl_NEVER = 0x0200
	gl_NICEST = 0x1102
	gl_NONE = 0
	gl_NOOP = 0x1505
	gl_NOR = 0x1508
	gl_NORMALIZE = 0x0BA1
	gl_NORMAL_ARRAY = 0x8075
	gl_NORMAL_ARRAY_BUFFER_BINDING = 0x8897
	gl_NORMAL_ARRAY_POINTER = 0x808F
	gl_NORMAL_ARRAY_STRIDE = 0x807F
	gl_NORMAL_ARRAY_TYPE = 0x807E
	gl_NORMAL_MAP = 0x8511
	gl_NOTEQUAL = 0x0205
	gl_NO_ERROR = 0
	gl_NUM_COMPRESSED_TEXTURE_FORMATS = 0x86A2
	gl_NUM_SHADER_BINARY_FORMATS = 0x8DF9
	gl_OBJECT_LINEAR = 0x2401
	gl_OBJECT_PLANE = 0x2501
	gl_ONE = 1
	gl_ONE_MINUS_CONSTANT_ALPHA = 0x8004
	gl_ONE_MINUS_CONSTANT_COLOR = 0x8002
	gl_ONE_MINUS_DST_ALPHA = 0x0305
	gl_ONE_MINUS_DST_COLOR = 0x0307
	gl_ONE_MINUS_SRC_ALPHA = 0x0303
	gl_ONE_MINUS_SRC_COLOR = 0x0301
	gl_OPERAND0_ALPHA = 0x8598
	gl_OPERAND0_RGB = 0x8590
	gl_OPERAND1_ALPHA = 0x8599
	gl_OPERAND1_RGB = 0x8591
	gl_OPERAND2_ALPHA = 0x859A
	gl_OPERAND2_RGB = 0x8592
	gl_OR = 0x1507
	gl_ORDER = 0x0A01
	gl_OR_INVERTED = 0x150D
	gl_OR_REVERSE = 0x150B
	gl_OUT_OF_MEMORY = 0x0505
	gl_PACK_ALIGNMENT = 0x0D05
	gl_PACK_IMAGE_HEIGHT = 0x806C
	gl_PACK_LSB_FIRST = 0x0D01
	gl_PACK_ROW_LENGTH = 0x0D02
	gl_PACK_SKIP_IMAGES = 0x806B
	gl_PACK_SKIP_PIXELS = 0x0D04
	gl_PACK_SKIP_ROWS = 0x0D03
	gl_PACK_SWAP_BYTES = 0x0D00
	gl_PASS_THROUGH_TOKEN = 0x0700
	gl_PERSPECTIVE_CORRECTION_HINT = 0x0C50
	gl_PIXEL_MAP_A_TO_A = 0x0C79
	gl_PIXEL_MAP_A_TO_A_SIZE = 0x0CB9
	gl_PIXEL_MAP_B_TO_B = 0x0C78
	gl_PIXEL_MAP_B_TO_B_SIZE = 0x0CB8
	gl_PIXEL_MAP_G_TO_G = 0x0C77
	gl_PIXEL_MAP_G_TO_G_SIZE = 0x0CB7
	gl_PIXEL_MAP_I_TO_A = 0x0C75
	gl_PIXEL_MAP_I_TO_A_SIZE = 0x0CB5
	gl_PIXEL_MAP_I_TO_B = 0x0C74
	gl_PIXEL_MAP_I_TO_B_SIZE = 0x0CB4
	gl_PIXEL_MAP_I_TO_G = 0x0C73
	gl_PIXEL_MAP_I_TO_G_SIZE = 0x0CB3
	gl_PIXEL_MAP_I_TO_I = 0x0C70
	gl_PIXEL_MAP_I_TO_I_SIZE = 0x0CB0
	gl_PIXEL_MAP_I_TO_R = 0x0C72
	gl_PIXEL_MAP_I_TO_R_SIZE = 0x0CB2
	gl_PIXEL_MAP_R_TO_R = 0x0C76
	gl_PIXEL_MAP_R_TO_R_SIZE = 0x0CB6
	gl_PIXEL_MAP_S_TO_S = 0x0C71
	gl_PIXEL_MAP_S_TO_S_SIZE = 0x0CB1
	gl_PIXEL_MODE_BIT = 0x00000020
	gl_POINT = 0x1B00
	gl_POINTS = 0x0000
	gl_POINT_BIT = 0x00000002
	gl_POINT_DISTANCE_ATTENUATION = 0x8129
	gl_POINT_FADE_THRESHOLD_SIZE = 0x8128
	gl_POINT_SIZE = 0x0B11
	gl_POINT_SIZE_GRANULARITY = 0x0B13
	gl_POINT_SIZE_MAX = 0x8127
	gl_POINT_SIZE_MIN = 0x8126
	gl_POINT_SIZE_RANGE = 0x0B12
	gl_POINT_SMOOTH = 0x0B10
	gl_POINT_SMOOTH_HINT = 0x0C51
	gl_POINT_SPRITE = 0x8861
	gl_POINT_SPRITE_COORD_ORIGIN = 0x8CA0
	gl_POINT_TOKEN = 0x0701
	gl_POLYGON = 0x0009
	gl_POLYGON_BIT = 0x00000008
	gl_POLYGON_MODE = 0x0B40
	gl_POLYGON_OFFSET_FACTOR = 0x8038
	gl_POLYGON_OFFSET_FILL = 0x8037
	gl_POLYGON_OFFSET_LINE = 0x2A02
	gl_POLYGON_OFFSET_POINT = 0x2A01
	gl_POLYGON_OFFSET_UNITS = 0x2A00
	gl_POLYGON_SMOOTH = 0x0B41
	gl_POLYGON_SMOOTH_HINT = 0x0C53
	gl_POLYGON_STIPPLE = 0x0B42
	gl_POLYGON_STIPPLE_BIT = 0x00000010
	gl_POLYGON_TOKEN = 0x0703
	gl_POSITION = 0x1203
	gl_PREVIOUS = 0x8578
	gl_PRIMARY_COLOR = 0x8577
	gl_PROJECTION = 0x1701
	gl_PROJECTION_MATRIX = 0x0BA7
	gl_PROJECTION_STACK_DEPTH = 0x0BA4
	gl_PROXY_TEXTURE_1D = 0x8063
	gl_PROXY_TEXTURE_2D = 0x8064
	gl_PROXY_TEXTURE_3D = 0x8070
	gl_PROXY_TEXTURE_CUBE_MAP = 0x851B
	gl_Q = 0x2003
	gl_QUADRATIC_ATTENUATION = 0x1209
	gl_QUADS = 0x0007
	gl_QUAD_STRIP = 0x0008
	gl_QUERY_COUNTER_BITS = 0x8864
	gl_QUERY_RESULT = 0x8866
	gl_QUERY_RESULT_AVAILABLE = 0x8867
	gl_R = 0x2002
	gl_R3_G3_B2 = 0x2A10
	gl_READ_BUFFER = 0x0C02
	gl_READ_ONLY = 0x88B8
	gl_READ_WRITE = 0x88BA
	gl_RED = 0x1903
	gl_RED_BIAS = 0x0D15
	gl_RED_BITS = 0x0D52
	gl_RED_SCALE = 0x0D14
	gl_REFLECTION_MAP = 0x8512
	gl_RENDER = 0x1C00
	gl_RENDERBUFFER = 0x8D41
	gl_RENDERBUFFER_ALPHA_SIZE = 0x8D53
	gl_RENDERBUFFER_BINDING = 0x8CA7
	gl_RENDERBUFFER_BLUE_SIZE = 0x8D52
	gl_RENDERBUFFER_DEPTH_SIZE = 0x8D54
	gl_RENDERBUFFER_GREEN_SIZE = 0x8D51
	gl_RENDERBUFFER_HEIGHT = 0x8D43
	gl_RENDERBUFFER_INTERNAL_FORMAT = 0x8D44
	gl_RENDERBUFFER_RED_SIZE = 0x8D50
	gl_RENDERBUFFER_STENCIL_SIZE = 0x8D55
	gl_RENDERBUFFER_WIDTH = 0x8D42
	gl_RENDERER = 0x1F01
	gl_RENDER_MODE = 0x0C40
	gl_REPEAT = 0x2901
	gl_REPLACE = 0x1E01
	gl_RESCALE_NORMAL = 0x803A
	gl_RETURN = 0x0102
	gl_RGB = 0x1907
	gl_RGB10 = 0x8052
	gl_RGB10_A2 = 0x8059
	gl_RGB12 = 0x8053
	gl_RGB16 = 0x8054
	gl_RGB4 = 0x804F
	gl_RGB5 = 0x8050
	gl_RGB565 = 0x8D62
	gl_RGB5_A1 = 0x8057
	gl_RGB8 = 0x8051
	gl_RGBA = 0x1908
	gl_RGBA12 = 0x805A
	gl_RGBA16 = 0x805B
	gl_RGBA2 = 0x8055
	gl_RGBA4 = 0x8056
	gl_RGBA8 = 0x8058
	gl_RGBA_MODE = 0x0C31
	gl_RGB_SCALE = 0x8573
	gl_RIGHT = 0x0407
	gl_S = 0x2000
	gl_SAMPLER_1D = 0x8B5D
	gl_SAMPLER_1D_SHADOW = 0x8B61
	gl_SAMPLER_2D = 0x8B5E
	gl_SAMPLER_2D_SHADOW = 0x8B62
	gl_SAMPLER_3D = 0x8B5F
	gl_SAMPLER_CUBE = 0x8B60
	gl_SAMPLES = 0x80A9
	gl_SAMPLES_PASSED = 0x8914
	gl_SAMPLE_ALPHA_TO_COVERAGE = 0x809E
	gl_SAMPLE_ALPHA_TO_ONE = 0x809F
	gl_SAMPLE_BUFFERS = 0x80A8
	gl_SAMPLE_COVERAGE = 0x80A0
	gl_SAMPLE_COVERAGE_INVERT = 0x80AB
	gl_SAMPLE_COVERAGE_VALUE = 0x80AA
	gl_SCISSOR_BIT = 0x00080000
	gl_SCISSOR_BOX = 0x0C10
	gl_SCISSOR_TEST = 0x0C11
	gl_SECONDARY_COLOR_ARRAY = 0x845E
	gl_SECONDARY_COLOR_ARRAY_BUFFER_BINDING = 0x889C
	gl_SECONDARY_COLOR_ARRAY_POINTER = 0x845D
	gl_SECONDARY_COLOR_ARRAY_SIZE = 0x845A
	gl_SECONDARY_COLOR_ARRAY_STRIDE = 0x845C
	gl_SECONDARY_COLOR_ARRAY_TYPE = 0x845B
	gl_SELECT = 0x1C02
	gl_SELECTION_BUFFER_POINTER = 0x0DF3
	gl_SELECTION_BUFFER_SIZE = 0x0DF4
	gl_SEPARATE_SPECULAR_COLOR = 0x81FA
	gl_SET = 0x150F
	gl_SHADER_BINARY_FORMATS = 0x8DF8
	gl_SHADER_COMPILER = 0x8DFA
	gl_SHADER_SOURCE_LENGTH = 0x8B88
	gl_SHADER_TYPE = 0x8B4F
	gl_SHADE_MODEL = 0x0B54
	gl_SHADING_LANGUAGE_VERSION = 0x8B8C
	gl_SHININESS = 0x1601
	gl_SHORT = 0x1402
	gl_SINGLE_COLOR = 0x81F9
	gl_SMOOTH = 0x1D01
	gl_SMOOTH_LINE_WIDTH_GRANULARITY = 0x0B23
	gl_SMOOTH_LINE_WIDTH_RANGE = 0x0B22
	gl_SMOOTH_POINT_SIZE_GRANULARITY = 0x0B13
	gl_SMOOTH_POINT_SIZE_RANGE = 0x0B12
	gl_SOURCE0_ALPHA = 0x8588
	gl_SOURCE0_RGB = 0x8580
	gl_SOURCE1_ALPHA = 0x8589
	gl_SOURCE1_RGB = 0x8581
	gl_SOURCE2_ALPHA = 0x858A
	gl_SOURCE2_RGB = 0x8582
	gl_SPECULAR = 0x1202
	gl_SPHERE_MAP = 0x2402
	gl_SPOT_CUTOFF = 0x1206
	gl_SPOT_DIRECTION = 0x1204
	gl_SPOT_EXPONENT = 0x1205
	gl_SRC0_ALPHA = 0x8588
	gl_SRC0_RGB = 0x8580
	gl_SRC1_ALPHA = 0x8589
	gl_SRC1_RGB = 0x8581
	gl_SRC2_ALPHA = 0x858A
	gl_SRC2_RGB = 0x8582
	gl_SRC_ALPHA = 0x0302
	gl_SRC_ALPHA_SATURATE = 0x0308
	gl_SRC_COLOR = 0x0300
	gl_STACK_OVERFLOW = 0x0503
	gl_STACK_UNDERFLOW = 0x0504
	gl_STATIC_COPY = 0x88E6
	gl_STATIC_DRAW = 0x88E4
	gl_STATIC_READ = 0x88E5
	gl_STENCIL = 0x1802
	gl_STENCIL_ATTACHMENT = 0x8D20
	gl_STENCIL_BACK_FAIL = 0x8801
	gl_STENCIL_BACK_FUNC = 0x8800
	gl_STENCIL_BACK_PASS_DEPTH_FAIL = 0x8802
	gl_STENCIL_BACK_PASS_DEPTH_PASS = 0x8803
	gl_STENCIL_BACK_REF = 0x8CA3
	gl_STENCIL_BACK_VALUE_MASK = 0x8CA4
	gl_STENCIL_BACK_WRITEMASK = 0x8CA5
	gl_STENCIL_BITS = 0x0D57
	gl_STENCIL_BUFFER_BIT = 0x00000400
	gl_STENCIL_CLEAR_VALUE = 0x0B91
	gl_STENCIL_FAIL = 0x0B94
	gl_STENCIL_FUNC = 0x0B92
	gl_STENCIL_INDEX = 0x1901
	gl_STENCIL_INDEX8 = 0x8D48
	gl_STENCIL_PASS_DEPTH_FAIL = 0x0B95
	gl_STENCIL_PASS_DEPTH_PASS = 0x0B96
	gl_STENCIL_REF = 0x0B97
	gl_STENCIL_TEST = 0x0B90
	gl_STENCIL_VALUE_MASK = 0x0B93
	gl_STENCIL_WRITEMASK = 0x0B98
	gl_STEREO = 0x0C33
	gl_STREAM_COPY = 0x88E2
	gl_STREAM_DRAW = 0x88E0
	gl_STREAM_READ = 0x88E1
	gl_SUBPIXEL_BITS = 0x0D50
	gl_SUBTRACT = 0x84E7
	gl_T = 0x2001
	gl_T2F_C3F_V3F = 0x2A2A
	gl_T2F_C4F_N3F_V3F = 0x2A2C
	gl_T2F_C4UB_V3F = 0x2A29
	gl_T2F_N3F_V3F = 0x2A2B
	gl_T2F_V3F = 0x2A27
	gl_T4F_C4F_N3F_V4F = 0x2A2D
	gl_T4F_V4F = 0x2A28
	gl_TEXTURE = 0x1702
	gl_TEXTURE0 = 0x84C0
	gl_TEXTURE1 = 0x84C1
	gl_TEXTURE10 = 0x84CA
	gl_TEXTURE11 = 0x84CB
	gl_TEXTURE12 = 0x84CC
	gl_TEXTURE13 = 0x84CD
	gl_TEXTURE14 = 0x84CE
	gl_TEXTURE15 = 0x84CF
	gl_TEXTURE16 = 0x84D0
	gl_TEXTURE17 = 0x84D1
	gl_TEXTURE18 = 0x84D2
	gl_TEXTURE19 = 0x84D3
	gl_TEXTURE2 = 0x84C2
	gl_TEXTURE20 = 0x84D4
	gl_TEXTURE21 = 0x84D5
	gl_TEXTURE22 = 0x84D6
	gl_TEXTURE23 = 0x84D7
	gl_TEXTURE24 = 0x84D8
	gl_TEXTURE25 = 0x84D9
	gl_TEXTURE26 = 0x84DA
	gl_TEXTURE27 = 0x84DB
	gl_TEXTURE28 = 0x84DC
	gl_TEXTURE29 = 0x84DD
	gl_TEXTURE3 = 0x84C3
	gl_TEXTURE30 = 0x84DE
	gl_TEXTURE31 = 0x84DF
	gl_TEXTURE4 = 0x84C4
	gl_TEXTURE5 = 0x84C5
	gl_TEXTURE6 = 0x84C6
	gl_TEXTURE7 = 0x84C7
	gl_TEXTURE8 = 0x84C8
	gl_TEXTURE9 = 0x84C9
	gl_TEXTURE_1D = 0x0DE0
	gl_TEXTURE_2D = 0x0DE1
	gl_TEXTURE_3D = 0x806F
	gl_TEXTURE_ALPHA_SIZE = 0x805F
	gl_TEXTURE_BASE_LEVEL = 0x813C
	gl_TEXTURE_BINDING_1D = 0x8068
	gl_TEXTURE_BINDING_2D = 0x8069
	gl_TEXTURE_BINDING_3D = 0x806A
	gl_TEXTURE_BINDING_CUBE_MAP = 0x8514
	gl_TEXTURE_BIT = 0x00040000
	gl_TEXTURE_BLUE_SIZE = 0x805E
	gl_TEXTURE_BORDER = 0x1005
	gl_TEXTURE_BORDER_COLOR = 0x1004
	gl_TEXTURE_COMPARE_FUNC = 0x884D
	gl_TEXTURE_COMPARE_MODE = 0x884C
	gl_TEXTURE_COMPONENTS = 0x1003
	gl_TEXTURE_COMPRESSED = 0x86A1
	gl_TEXTURE_COMPRESSED_IMAGE_SIZE = 0x86A0
	gl_TEXTURE_COMPRESSION_HINT = 0x84EF
	gl_TEXTURE_COORD_ARRAY = 0x8078
	gl_TEXTURE_COORD_ARRAY_BUFFER_BINDING = 0x889A
	gl_TEXTURE_COORD_ARRAY_POINTER = 0x8092
	gl_TEXTURE_COORD_ARRAY_SIZE = 0x8088
	gl_TEXTURE_COORD_ARRAY_STRIDE = 0x808A
	gl_TEXTURE_COORD_ARRAY_TYPE = 0x8089
	gl_TEXTURE_CUBE_MAP = 0x8513
	gl_TEXTURE_CUBE_MAP_NEGATIVE_X = 0x8516
	gl_TEXTURE_CUBE_MAP_NEGATIVE_Y = 0x8518
	gl_TEXTURE_CUBE_MAP_NEGATIVE_Z = 0x851A
	gl_TEXTURE_CUBE_MAP_POSITIVE_X = 0x8515
	gl_TEXTURE_CUBE_MAP_POSITIVE_Y = 0x8517
	gl_TEXTURE_CUBE_MAP_POSITIVE_Z = 0x8519
	gl_TEXTURE_DEPTH = 0x8071
	gl_TEXTURE_DEPTH_SIZE = 0x884A
	gl_TEXTURE_ENV = 0x2300
	gl_TEXTURE_ENV_COLOR = 0x2201
	gl_TEXTURE_ENV_MODE = 0x2200
	gl_TEXTURE_FILTER_CONTROL = 0x8500
	gl_TEXTURE_GEN_MODE = 0x2500
	gl_TEXTURE_GEN_Q = 0x0C63
	gl_TEXTURE_GEN_R = 0x0C62
	gl_TEXTURE_GEN_S = 0x0C60
	gl_TEXTURE_GEN_T = 0x0C61
	gl_TEXTURE_GREEN_SIZE = 0x805D
	gl_TEXTURE_HEIGHT = 0x1001
	gl_TEXTURE_INTENSITY_SIZE = 0x8061
	gl_TEXTURE_INTERNAL_FORMAT = 0x1003
	gl_TEXTURE_LOD_BIAS = 0x8501
	gl_TEXTURE_LUMINANCE_SIZE = 0x8060
	gl_TEXTURE_MAG_FILTER = 0x2800
	gl_TEXTURE_MATRIX = 0x0BA8
	gl_TEXTURE_MAX_LEVEL = 0x813D
	gl_TEXTURE_MAX_LOD = 0x813B
	gl_TEXTURE_MIN_FILTER = 0x2801
	gl_TEXTURE_MIN_LOD = 0x813A
	gl_TEXTURE_PRIORITY = 0x8066
	gl_TEXTURE_RED_SIZE = 0x805C
	gl_TEXTURE_RESIDENT = 0x8067
	gl_TEXTURE_STACK_DEPTH = 0x0BA5
	gl_TEXTURE_WIDTH = 0x1000
	gl_TEXTURE_WRAP_R = 0x8072
	gl_TEXTURE_WRAP_S = 0x2802
	gl_TEXTURE_WRAP_T = 0x2803
	gl_TRANSFORM_BIT = 0x00001000
	gl_TRANSPOSE_COLOR_MATRIX = 0x84E6
	gl_TRANSPOSE_MODELVIEW_MATRIX = 0x84E3
	gl_TRANSPOSE_PROJECTION_MATRIX = 0x84E4
	gl_TRANSPOSE_TEXTURE_MATRIX = 0x84E5
	gl_TRIANGLES = 0x0004
	gl_TRIANGLE_FAN = 0x0006
	gl_TRIANGLE_STRIP = 0x0005
	gl_TRUE = 1
	gl_UNPACK_ALIGNMENT = 0x0CF5
	gl_UNPACK_IMAGE_HEIGHT = 0x806E
	gl_UNPACK_LSB_FIRST = 0x0CF1
	gl_UNPACK_ROW_LENGTH = 0x0CF2
	gl_UNPACK_SKIP_IMAGES = 0x806D
	gl_UNPACK_SKIP_PIXELS = 0x0CF4
	gl_UNPACK_SKIP_ROWS = 0x0CF3
	gl_UNPACK_SWAP_BYTES = 0x0CF0
	gl_UNSIGNED_BYTE = 0x1401
	gl_UNSIGNED_BYTE_2_3_3_REV = 0x8362
	gl_UNSIGNED_BYTE_3_3_2 = 0x8032
	gl_UNSIGNED_INT = 0x1405
	gl_UNSIGNED_INT_10_10_10_2 = 0x8036
	gl_UNSIGNED_INT_2_10_10_10_REV = 0x8368
	gl_UNSIGNED_INT_8_8_8_8 = 0x8035
	gl_UNSIGNED_INT_8_8_8_8_REV = 0x8367
	gl_UNSIGNED_SHORT = 0x1403
	gl_UNSIGNED_SHORT_1_5_5_5_REV = 0x8366
	gl_UNSIGNED_SHORT_4_4_4_4 = 0x8033
	gl_UNSIGNED_SHORT_4_4_4_4_REV = 0x8365
	gl_UNSIGNED_SHORT_5_5_5_1 = 0x8034
	gl_UNSIGNED_SHORT_5_6_5 = 0x8363
	gl_UNSIGNED_SHORT_5_6_5_REV = 0x8364
	gl_UPPER_LEFT = 0x8CA2
	gl_V2F = 0x2A20
	gl_V3F = 0x2A21
	gl_VALIDATE_STATUS = 0x8B83
	gl_VENDOR = 0x1F00
	gl_VERSION = 0x1F02
	gl_VERSION_ES_CL_1_0 = 1
	gl_VERSION_ES_CL_1_1 = 1
	gl_VERSION_ES_CM_1_1 = 1
	gl_VERTEX_ARRAY = 0x8074
	gl_VERTEX_ARRAY_BUFFER_BINDING = 0x8896
	gl_VERTEX_ARRAY_POINTER = 0x808E
	gl_VERTEX_ARRAY_SIZE = 0x807A
	gl_VERTEX_ARRAY_STRIDE = 0x807C
	gl_VERTEX_ARRAY_TYPE = 0x807B
	gl_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING = 0x889F
	gl_VERTEX_ATTRIB_ARRAY_ENABLED = 0x8622
	gl_VERTEX_ATTRIB_ARRAY_NORMALIZED = 0x886A
	gl_VERTEX_ATTRIB_ARRAY_POINTER = 0x8645
	gl_VERTEX_ATTRIB_ARRAY_SIZE = 0x8623
	gl_VERTEX_ATTRIB_ARRAY_STRIDE = 0x8624
	gl_VERTEX_ATTRIB_ARRAY_TYPE = 0x8625
	gl_VERTEX_PROGRAM_POINT_SIZE = 0x8642
	gl_VERTEX_PROGRAM_TWO_SIDE = 0x8643
	gl_VERTEX_SHADER = 0x8B31
	gl_VIEWPORT = 0x0BA2
	gl_VIEWPORT_BIT = 0x00000800
	gl_WEIGHT_ARRAY_BUFFER_BINDING = 0x889E
	gl_WRITE_ONLY = 0x88B9
	gl_XOR = 0x1506
	gl_ZERO = 0
	gl_ZOOM_X = 0x0D16
	gl_ZOOM_Y = 0x0D17
)
// package gl EOF
