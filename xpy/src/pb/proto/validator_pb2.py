# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/validator.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import descriptor_pb2 as google_dot_protobuf_dot_descriptor__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x15proto/validator.proto\x12\tvalidator\x1a google/protobuf/descriptor.proto\"Y\n\x0cValidOptions\x12$\n\x08\x63heck_if\x18\x01 \x01(\x0b\x32\x12.validator.CheckIf\x12#\n\x04tags\x18\x02 \x01(\x0b\x32\x15.validator.TagOptions\"=\n\x07\x43heckIf\x12\r\n\x05\x66ield\x18\x01 \x01(\t\x12#\n\x04tags\x18\x02 \x01(\x0b\x32\x15.validator.TagOptions\"\xbf\x03\n\nTagOptions\x12%\n\x05oneof\x18\x14 \x01(\x0b\x32\x14.validator.OneOfTagsH\x00\x12%\n\x05\x66loat\x18\x15 \x01(\x0b\x32\x14.validator.FloatTagsH\x00\x12!\n\x03int\x18\x16 \x01(\x0b\x32\x12.validator.IntTagsH\x00\x12#\n\x04uint\x18\x17 \x01(\x0b\x32\x13.validator.UintTagsH\x00\x12\'\n\x06string\x18\x18 \x01(\x0b\x32\x15.validator.StringTagsH\x00\x12%\n\x05\x62ytes\x18\x19 \x01(\x0b\x32\x14.validator.BytesTagsH\x00\x12#\n\x04\x62ool\x18\x1a \x01(\x0b\x32\x13.validator.BoolTagsH\x00\x12#\n\x04\x65num\x18\x1b \x01(\x0b\x32\x13.validator.EnumTagsH\x00\x12)\n\x07message\x18\x1c \x01(\x0b\x32\x16.validator.MessageTagsH\x00\x12+\n\x08repeated\x18\x1d \x01(\x0b\x32\x17.validator.RepeatedTagsH\x00\x12!\n\x03map\x18\x1e \x01(\x0b\x32\x12.validator.MapTagsH\x00\x42\x06\n\x04kind\"/\n\tOneOfTags\x12\x15\n\x08not_null\x18\x01 \x01(\x08H\x00\x88\x01\x01\x42\x0b\n\t_not_null\"\xbb\x01\n\tFloatTags\x12\x0f\n\x02\x65q\x18\x03 \x01(\x01H\x00\x88\x01\x01\x12\x0f\n\x02ne\x18\x04 \x01(\x01H\x01\x88\x01\x01\x12\x0f\n\x02lt\x18\x05 \x01(\x01H\x02\x88\x01\x01\x12\x0f\n\x02gt\x18\x06 \x01(\x01H\x03\x88\x01\x01\x12\x10\n\x03lte\x18\x07 \x01(\x01H\x04\x88\x01\x01\x12\x10\n\x03gte\x18\x08 \x01(\x01H\x05\x88\x01\x01\x12\n\n\x02in\x18\t \x03(\x01\x12\x0e\n\x06not_in\x18\n \x03(\x01\x42\x05\n\x03_eqB\x05\n\x03_neB\x05\n\x03_ltB\x05\n\x03_gtB\x06\n\x04_lteB\x06\n\x04_gte\"\xb9\x01\n\x07IntTags\x12\x0f\n\x02\x65q\x18\x03 \x01(\x03H\x00\x88\x01\x01\x12\x0f\n\x02ne\x18\x04 \x01(\x03H\x01\x88\x01\x01\x12\x0f\n\x02lt\x18\x05 \x01(\x03H\x02\x88\x01\x01\x12\x0f\n\x02gt\x18\x06 \x01(\x03H\x03\x88\x01\x01\x12\x10\n\x03lte\x18\x07 \x01(\x03H\x04\x88\x01\x01\x12\x10\n\x03gte\x18\x08 \x01(\x03H\x05\x88\x01\x01\x12\n\n\x02in\x18\t \x03(\x03\x12\x0e\n\x06not_in\x18\n \x03(\x03\x42\x05\n\x03_eqB\x05\n\x03_neB\x05\n\x03_ltB\x05\n\x03_gtB\x06\n\x04_lteB\x06\n\x04_gte\"\xba\x01\n\x08UintTags\x12\x0f\n\x02\x65q\x18\x03 \x01(\x04H\x00\x88\x01\x01\x12\x0f\n\x02ne\x18\x04 \x01(\x04H\x01\x88\x01\x01\x12\x0f\n\x02lt\x18\x05 \x01(\x04H\x02\x88\x01\x01\x12\x0f\n\x02gt\x18\x06 \x01(\x04H\x03\x88\x01\x01\x12\x10\n\x03lte\x18\x07 \x01(\x04H\x04\x88\x01\x01\x12\x10\n\x03gte\x18\x08 \x01(\x04H\x05\x88\x01\x01\x12\n\n\x02in\x18\t \x03(\x04\x12\x0e\n\x06not_in\x18\n \x03(\x04\x42\x05\n\x03_eqB\x05\n\x03_neB\x05\n\x03_ltB\x05\n\x03_gtB\x06\n\x04_lteB\x06\n\x04_gte\"\xcb\x15\n\nStringTags\x12\x0f\n\x02\x65q\x18\x03 \x01(\tH\x00\x88\x01\x01\x12\x0f\n\x02ne\x18\x04 \x01(\tH\x01\x88\x01\x01\x12\x0f\n\x02lt\x18\x05 \x01(\tH\x02\x88\x01\x01\x12\x0f\n\x02gt\x18\x06 \x01(\tH\x03\x88\x01\x01\x12\x10\n\x03lte\x18\x07 \x01(\tH\x04\x88\x01\x01\x12\x10\n\x03gte\x18\x08 \x01(\tH\x05\x88\x01\x01\x12\n\n\x02in\x18\t \x03(\t\x12\x0e\n\x06not_in\x18\n \x03(\t\x12\x18\n\x0b\x63har_len_eq\x18\x14 \x01(\x03H\x06\x88\x01\x01\x12\x18\n\x0b\x63har_len_ne\x18\x15 \x01(\x03H\x07\x88\x01\x01\x12\x18\n\x0b\x63har_len_gt\x18\x16 \x01(\x03H\x08\x88\x01\x01\x12\x18\n\x0b\x63har_len_lt\x18\x17 \x01(\x03H\t\x88\x01\x01\x12\x19\n\x0c\x63har_len_gte\x18\x18 \x01(\x03H\n\x88\x01\x01\x12\x19\n\x0c\x63har_len_lte\x18\x19 \x01(\x03H\x0b\x88\x01\x01\x12\x18\n\x0b\x62yte_len_eq\x18\x1e \x01(\x03H\x0c\x88\x01\x01\x12\x18\n\x0b\x62yte_len_ne\x18\x1f \x01(\x03H\r\x88\x01\x01\x12\x18\n\x0b\x62yte_len_gt\x18  \x01(\x03H\x0e\x88\x01\x01\x12\x18\n\x0b\x62yte_len_lt\x18! \x01(\x03H\x0f\x88\x01\x01\x12\x19\n\x0c\x62yte_len_gte\x18\" \x01(\x03H\x10\x88\x01\x01\x12\x19\n\x0c\x62yte_len_lte\x18# \x01(\x03H\x11\x88\x01\x01\x12\x12\n\x05regex\x18( \x01(\tH\x12\x88\x01\x01\x12\x13\n\x06prefix\x18) \x01(\tH\x13\x88\x01\x01\x12\x16\n\tno_prefix\x18* \x01(\tH\x14\x88\x01\x01\x12\x13\n\x06suffix\x18+ \x01(\tH\x15\x88\x01\x01\x12\x16\n\tno_suffix\x18, \x01(\tH\x16\x88\x01\x01\x12\x15\n\x08\x63ontains\x18- \x01(\tH\x17\x88\x01\x01\x12\x19\n\x0cnot_contains\x18. \x01(\tH\x18\x88\x01\x01\x12\x19\n\x0c\x63ontains_any\x18/ \x01(\tH\x19\x88\x01\x01\x12\x1d\n\x10not_contains_any\x18\x30 \x01(\tH\x1a\x88\x01\x01\x12\x11\n\x04utf8\x18Q \x01(\x08H\x1b\x88\x01\x01\x12\x12\n\x05\x61scii\x18G \x01(\x08H\x1c\x88\x01\x01\x12\x18\n\x0bprint_ascii\x18H \x01(\x08H\x1d\x88\x01\x01\x12\x14\n\x07\x62oolean\x18I \x01(\x08H\x1e\x88\x01\x01\x12\x16\n\tlowercase\x18J \x01(\x08H\x1f\x88\x01\x01\x12\x16\n\tuppercase\x18K \x01(\x08H \x88\x01\x01\x12\x12\n\x05\x61lpha\x18L \x01(\x08H!\x88\x01\x01\x12\x13\n\x06number\x18M \x01(\x08H\"\x88\x01\x01\x12\x19\n\x0c\x61lpha_number\x18N \x01(\x08H#\x88\x01\x01\x12\x0f\n\x02ip\x18\x65 \x01(\x08H$\x88\x01\x01\x12\x11\n\x04ipv4\x18\x66 \x01(\x08H%\x88\x01\x01\x12\x11\n\x04ipv6\x18g \x01(\x08H&\x88\x01\x01\x12\x14\n\x07ip_addr\x18h \x01(\x08H\'\x88\x01\x01\x12\x15\n\x08ip4_addr\x18i \x01(\x08H(\x88\x01\x01\x12\x15\n\x08ip6_addr\x18j \x01(\x08H)\x88\x01\x01\x12\x11\n\x04\x63idr\x18k \x01(\x08H*\x88\x01\x01\x12\x13\n\x06\x63idrv4\x18l \x01(\x08H+\x88\x01\x01\x12\x13\n\x06\x63idrv6\x18m \x01(\x08H,\x88\x01\x01\x12\x15\n\x08tcp_addr\x18o \x01(\x08H-\x88\x01\x01\x12\x16\n\ttcp4_addr\x18p \x01(\x08H.\x88\x01\x01\x12\x16\n\ttcp6_addr\x18q \x01(\x08H/\x88\x01\x01\x12\x15\n\x08udp_addr\x18r \x01(\x08H0\x88\x01\x01\x12\x16\n\tudp4_addr\x18s \x01(\x08H1\x88\x01\x01\x12\x16\n\tudp6_addr\x18t \x01(\x08H2\x88\x01\x01\x12\x10\n\x03mac\x18n \x01(\x08H3\x88\x01\x01\x12\x16\n\tunix_addr\x18u \x01(\x08H4\x88\x01\x01\x12\x15\n\x08hostname\x18v \x01(\x08H5\x88\x01\x01\x12\x1d\n\x10hostname_rfc1123\x18w \x01(\x08H6\x88\x01\x01\x12\x1a\n\rhostname_port\x18x \x01(\x08H7\x88\x01\x01\x12\x15\n\x08\x64\x61ta_uri\x18y \x01(\x08H8\x88\x01\x01\x12\x11\n\x04\x66qdn\x18z \x01(\x08H9\x88\x01\x01\x12\x10\n\x03uri\x18{ \x01(\x08H:\x88\x01\x01\x12\x10\n\x03url\x18| \x01(\x08H;\x88\x01\x01\x12\x18\n\x0burl_encoded\x18} \x01(\x08H<\x88\x01\x01\x12\x16\n\tunix_cron\x18P \x01(\x08H=\x88\x01\x01\x12\x13\n\x05\x65mail\x18\x8c\x01 \x01(\x08H>\x88\x01\x01\x12\x12\n\x04json\x18\x8d\x01 \x01(\x08H?\x88\x01\x01\x12\x11\n\x03jwt\x18\x8e\x01 \x01(\x08H@\x88\x01\x01\x12\x12\n\x04html\x18\x8f\x01 \x01(\x08HA\x88\x01\x01\x12\x1a\n\x0chtml_encoded\x18\x90\x01 \x01(\x08HB\x88\x01\x01\x12\x14\n\x06\x62\x61se64\x18\x91\x01 \x01(\x08HC\x88\x01\x01\x12\x18\n\nbase64_url\x18\x92\x01 \x01(\x08HD\x88\x01\x01\x12\x19\n\x0bhexadecimal\x18\x93\x01 \x01(\x08HE\x88\x01\x01\x12\x16\n\x08\x64\x61tetime\x18\x94\x01 \x01(\tHF\x88\x01\x01\x12\x16\n\x08timezone\x18\x95\x01 \x01(\x08HG\x88\x01\x01\x12\x12\n\x04uuid\x18\x96\x01 \x01(\x08HH\x88\x01\x01\x12\x13\n\x05uuid1\x18\x97\x01 \x01(\x08HI\x88\x01\x01\x12\x13\n\x05uuid3\x18\x98\x01 \x01(\x08HJ\x88\x01\x01\x12\x13\n\x05uuid4\x18\x99\x01 \x01(\x08HK\x88\x01\x01\x12\x13\n\x05uuid5\x18\x9a\x01 \x01(\x08HL\x88\x01\x01\x42\x05\n\x03_eqB\x05\n\x03_neB\x05\n\x03_ltB\x05\n\x03_gtB\x06\n\x04_lteB\x06\n\x04_gteB\x0e\n\x0c_char_len_eqB\x0e\n\x0c_char_len_neB\x0e\n\x0c_char_len_gtB\x0e\n\x0c_char_len_ltB\x0f\n\r_char_len_gteB\x0f\n\r_char_len_lteB\x0e\n\x0c_byte_len_eqB\x0e\n\x0c_byte_len_neB\x0e\n\x0c_byte_len_gtB\x0e\n\x0c_byte_len_ltB\x0f\n\r_byte_len_gteB\x0f\n\r_byte_len_lteB\x08\n\x06_regexB\t\n\x07_prefixB\x0c\n\n_no_prefixB\t\n\x07_suffixB\x0c\n\n_no_suffixB\x0b\n\t_containsB\x0f\n\r_not_containsB\x0f\n\r_contains_anyB\x13\n\x11_not_contains_anyB\x07\n\x05_utf8B\x08\n\x06_asciiB\x0e\n\x0c_print_asciiB\n\n\x08_booleanB\x0c\n\n_lowercaseB\x0c\n\n_uppercaseB\x08\n\x06_alphaB\t\n\x07_numberB\x0f\n\r_alpha_numberB\x05\n\x03_ipB\x07\n\x05_ipv4B\x07\n\x05_ipv6B\n\n\x08_ip_addrB\x0b\n\t_ip4_addrB\x0b\n\t_ip6_addrB\x07\n\x05_cidrB\t\n\x07_cidrv4B\t\n\x07_cidrv6B\x0b\n\t_tcp_addrB\x0c\n\n_tcp4_addrB\x0c\n\n_tcp6_addrB\x0b\n\t_udp_addrB\x0c\n\n_udp4_addrB\x0c\n\n_udp6_addrB\x06\n\x04_macB\x0c\n\n_unix_addrB\x0b\n\t_hostnameB\x13\n\x11_hostname_rfc1123B\x10\n\x0e_hostname_portB\x0b\n\t_data_uriB\x07\n\x05_fqdnB\x06\n\x04_uriB\x06\n\x04_urlB\x0e\n\x0c_url_encodedB\x0c\n\n_unix_cronB\x08\n\x06_emailB\x07\n\x05_jsonB\x06\n\x04_jwtB\x07\n\x05_htmlB\x0f\n\r_html_encodedB\t\n\x07_base64B\r\n\x0b_base64_urlB\x0e\n\x0c_hexadecimalB\x0b\n\t_datetimeB\x0b\n\t_timezoneB\x07\n\x05_uuidB\x08\n\x06_uuid1B\x08\n\x06_uuid3B\x08\n\x06_uuid4B\x08\n\x06_uuid5\"\xcf\x01\n\tBytesTags\x12\x13\n\x06len_eq\x18\x03 \x01(\x03H\x00\x88\x01\x01\x12\x13\n\x06len_ne\x18\x04 \x01(\x03H\x01\x88\x01\x01\x12\x13\n\x06len_lt\x18\x05 \x01(\x03H\x02\x88\x01\x01\x12\x13\n\x06len_gt\x18\x06 \x01(\x03H\x03\x88\x01\x01\x12\x14\n\x07len_lte\x18\x07 \x01(\x03H\x04\x88\x01\x01\x12\x14\n\x07len_gte\x18\x08 \x01(\x03H\x05\x88\x01\x01\x42\t\n\x07_len_eqB\t\n\x07_len_neB\t\n\x07_len_ltB\t\n\x07_len_gtB\n\n\x08_len_lteB\n\n\x08_len_gte\"\"\n\x08\x42oolTags\x12\x0f\n\x02\x65q\x18\x03 \x01(\x08H\x00\x88\x01\x01\x42\x05\n\x03_eq\"\xde\x01\n\x08\x45numTags\x12\x0f\n\x02\x65q\x18\x03 \x01(\x05H\x00\x88\x01\x01\x12\x0f\n\x02ne\x18\x04 \x01(\x05H\x01\x88\x01\x01\x12\x0f\n\x02lt\x18\x05 \x01(\x05H\x02\x88\x01\x01\x12\x0f\n\x02gt\x18\x06 \x01(\x05H\x03\x88\x01\x01\x12\x10\n\x03lte\x18\x07 \x01(\x05H\x04\x88\x01\x01\x12\x10\n\x03gte\x18\x08 \x01(\x05H\x05\x88\x01\x01\x12\n\n\x02in\x18\t \x03(\x05\x12\x0e\n\x06not_in\x18\n \x03(\x05\x12\x15\n\x08in_enums\x18\x0b \x01(\x08H\x06\x88\x01\x01\x42\x05\n\x03_eqB\x05\n\x03_neB\x05\n\x03_ltB\x05\n\x03_gtB\x06\n\x04_lteB\x06\n\x04_gteB\x0b\n\t_in_enums\"M\n\x0bMessageTags\x12\x15\n\x08not_null\x18\x02 \x01(\x08H\x00\x88\x01\x01\x12\x11\n\x04skip\x18\x03 \x01(\x08H\x01\x88\x01\x01\x42\x0b\n\t_not_nullB\x07\n\x05_skip\"\xbb\x02\n\x0cRepeatedTags\x12\x15\n\x08not_null\x18\x02 \x01(\x08H\x00\x88\x01\x01\x12\x13\n\x06len_eq\x18\x03 \x01(\x03H\x01\x88\x01\x01\x12\x13\n\x06len_ne\x18\x04 \x01(\x03H\x02\x88\x01\x01\x12\x13\n\x06len_lt\x18\x05 \x01(\x03H\x03\x88\x01\x01\x12\x13\n\x06len_gt\x18\x06 \x01(\x03H\x04\x88\x01\x01\x12\x14\n\x07len_lte\x18\x07 \x01(\x03H\x05\x88\x01\x01\x12\x14\n\x07len_gte\x18\x08 \x01(\x03H\x06\x88\x01\x01\x12\x13\n\x06unique\x18\n \x01(\x08H\x07\x88\x01\x01\x12#\n\x04item\x18\x0b \x01(\x0b\x32\x15.validator.TagOptionsB\x0b\n\t_not_nullB\t\n\x07_len_eqB\t\n\x07_len_neB\t\n\x07_len_ltB\t\n\x07_len_gtB\n\n\x08_len_lteB\n\n\x08_len_gteB\t\n\x07_unique\"\xbb\x02\n\x07MapTags\x12\x15\n\x08not_null\x18\x02 \x01(\x08H\x00\x88\x01\x01\x12\x13\n\x06len_eq\x18\x03 \x01(\x03H\x01\x88\x01\x01\x12\x13\n\x06len_ne\x18\x04 \x01(\x03H\x02\x88\x01\x01\x12\x13\n\x06len_lt\x18\x05 \x01(\x03H\x03\x88\x01\x01\x12\x13\n\x06len_gt\x18\x06 \x01(\x03H\x04\x88\x01\x01\x12\x14\n\x07len_lte\x18\x07 \x01(\x03H\x05\x88\x01\x01\x12\x14\n\x07len_gte\x18\x08 \x01(\x03H\x06\x88\x01\x01\x12\"\n\x03key\x18\x0b \x01(\x0b\x32\x15.validator.TagOptions\x12$\n\x05value\x18\x0c \x01(\x0b\x32\x15.validator.TagOptionsB\x0b\n\t_not_nullB\t\n\x07_len_eqB\t\n\x07_len_neB\t\n\x07_len_ltB\t\n\x07_len_gtB\n\n\x08_len_lteB\n\n\x08_len_gte:G\n\x05\x66ield\x12\x1d.google.protobuf.FieldOptions\x18\xfc\xfb\x03 \x01(\x0b\x32\x17.validator.ValidOptions:G\n\x05oneof\x12\x1d.google.protobuf.OneofOptions\x18\x87\xfc\x03 \x01(\x0b\x32\x17.validator.ValidOptionsBg\n$io.github.yu31.protoc.pb.pbvalidatorB\x0bPBValidatorP\x00Z0github.com/yu31/protoc-plugin/xgo/pb/pbvalidatorb\x06proto3')


FIELD_FIELD_NUMBER = 65020
field = DESCRIPTOR.extensions_by_name['field']
ONEOF_FIELD_NUMBER = 65031
oneof = DESCRIPTOR.extensions_by_name['oneof']

_VALIDOPTIONS = DESCRIPTOR.message_types_by_name['ValidOptions']
_CHECKIF = DESCRIPTOR.message_types_by_name['CheckIf']
_TAGOPTIONS = DESCRIPTOR.message_types_by_name['TagOptions']
_ONEOFTAGS = DESCRIPTOR.message_types_by_name['OneOfTags']
_FLOATTAGS = DESCRIPTOR.message_types_by_name['FloatTags']
_INTTAGS = DESCRIPTOR.message_types_by_name['IntTags']
_UINTTAGS = DESCRIPTOR.message_types_by_name['UintTags']
_STRINGTAGS = DESCRIPTOR.message_types_by_name['StringTags']
_BYTESTAGS = DESCRIPTOR.message_types_by_name['BytesTags']
_BOOLTAGS = DESCRIPTOR.message_types_by_name['BoolTags']
_ENUMTAGS = DESCRIPTOR.message_types_by_name['EnumTags']
_MESSAGETAGS = DESCRIPTOR.message_types_by_name['MessageTags']
_REPEATEDTAGS = DESCRIPTOR.message_types_by_name['RepeatedTags']
_MAPTAGS = DESCRIPTOR.message_types_by_name['MapTags']
ValidOptions = _reflection.GeneratedProtocolMessageType('ValidOptions', (_message.Message,), {
  'DESCRIPTOR' : _VALIDOPTIONS,
  '__module__' : 'proto.validator_pb2'
  # @@protoc_insertion_point(class_scope:validator.ValidOptions)
  })
_sym_db.RegisterMessage(ValidOptions)

CheckIf = _reflection.GeneratedProtocolMessageType('CheckIf', (_message.Message,), {
  'DESCRIPTOR' : _CHECKIF,
  '__module__' : 'proto.validator_pb2'
  # @@protoc_insertion_point(class_scope:validator.CheckIf)
  })
_sym_db.RegisterMessage(CheckIf)

TagOptions = _reflection.GeneratedProtocolMessageType('TagOptions', (_message.Message,), {
  'DESCRIPTOR' : _TAGOPTIONS,
  '__module__' : 'proto.validator_pb2'
  # @@protoc_insertion_point(class_scope:validator.TagOptions)
  })
_sym_db.RegisterMessage(TagOptions)

OneOfTags = _reflection.GeneratedProtocolMessageType('OneOfTags', (_message.Message,), {
  'DESCRIPTOR' : _ONEOFTAGS,
  '__module__' : 'proto.validator_pb2'
  # @@protoc_insertion_point(class_scope:validator.OneOfTags)
  })
_sym_db.RegisterMessage(OneOfTags)

FloatTags = _reflection.GeneratedProtocolMessageType('FloatTags', (_message.Message,), {
  'DESCRIPTOR' : _FLOATTAGS,
  '__module__' : 'proto.validator_pb2'
  # @@protoc_insertion_point(class_scope:validator.FloatTags)
  })
_sym_db.RegisterMessage(FloatTags)

IntTags = _reflection.GeneratedProtocolMessageType('IntTags', (_message.Message,), {
  'DESCRIPTOR' : _INTTAGS,
  '__module__' : 'proto.validator_pb2'
  # @@protoc_insertion_point(class_scope:validator.IntTags)
  })
_sym_db.RegisterMessage(IntTags)

UintTags = _reflection.GeneratedProtocolMessageType('UintTags', (_message.Message,), {
  'DESCRIPTOR' : _UINTTAGS,
  '__module__' : 'proto.validator_pb2'
  # @@protoc_insertion_point(class_scope:validator.UintTags)
  })
_sym_db.RegisterMessage(UintTags)

StringTags = _reflection.GeneratedProtocolMessageType('StringTags', (_message.Message,), {
  'DESCRIPTOR' : _STRINGTAGS,
  '__module__' : 'proto.validator_pb2'
  # @@protoc_insertion_point(class_scope:validator.StringTags)
  })
_sym_db.RegisterMessage(StringTags)

BytesTags = _reflection.GeneratedProtocolMessageType('BytesTags', (_message.Message,), {
  'DESCRIPTOR' : _BYTESTAGS,
  '__module__' : 'proto.validator_pb2'
  # @@protoc_insertion_point(class_scope:validator.BytesTags)
  })
_sym_db.RegisterMessage(BytesTags)

BoolTags = _reflection.GeneratedProtocolMessageType('BoolTags', (_message.Message,), {
  'DESCRIPTOR' : _BOOLTAGS,
  '__module__' : 'proto.validator_pb2'
  # @@protoc_insertion_point(class_scope:validator.BoolTags)
  })
_sym_db.RegisterMessage(BoolTags)

EnumTags = _reflection.GeneratedProtocolMessageType('EnumTags', (_message.Message,), {
  'DESCRIPTOR' : _ENUMTAGS,
  '__module__' : 'proto.validator_pb2'
  # @@protoc_insertion_point(class_scope:validator.EnumTags)
  })
_sym_db.RegisterMessage(EnumTags)

MessageTags = _reflection.GeneratedProtocolMessageType('MessageTags', (_message.Message,), {
  'DESCRIPTOR' : _MESSAGETAGS,
  '__module__' : 'proto.validator_pb2'
  # @@protoc_insertion_point(class_scope:validator.MessageTags)
  })
_sym_db.RegisterMessage(MessageTags)

RepeatedTags = _reflection.GeneratedProtocolMessageType('RepeatedTags', (_message.Message,), {
  'DESCRIPTOR' : _REPEATEDTAGS,
  '__module__' : 'proto.validator_pb2'
  # @@protoc_insertion_point(class_scope:validator.RepeatedTags)
  })
_sym_db.RegisterMessage(RepeatedTags)

MapTags = _reflection.GeneratedProtocolMessageType('MapTags', (_message.Message,), {
  'DESCRIPTOR' : _MAPTAGS,
  '__module__' : 'proto.validator_pb2'
  # @@protoc_insertion_point(class_scope:validator.MapTags)
  })
_sym_db.RegisterMessage(MapTags)

if _descriptor._USE_C_DESCRIPTORS == False:
  google_dot_protobuf_dot_descriptor__pb2.FieldOptions.RegisterExtension(field)
  google_dot_protobuf_dot_descriptor__pb2.OneofOptions.RegisterExtension(oneof)

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n$io.github.yu31.protoc.pb.pbvalidatorB\013PBValidatorP\000Z0github.com/yu31/protoc-plugin/xgo/pb/pbvalidator'
  _VALIDOPTIONS._serialized_start=70
  _VALIDOPTIONS._serialized_end=159
  _CHECKIF._serialized_start=161
  _CHECKIF._serialized_end=222
  _TAGOPTIONS._serialized_start=225
  _TAGOPTIONS._serialized_end=672
  _ONEOFTAGS._serialized_start=674
  _ONEOFTAGS._serialized_end=721
  _FLOATTAGS._serialized_start=724
  _FLOATTAGS._serialized_end=911
  _INTTAGS._serialized_start=914
  _INTTAGS._serialized_end=1099
  _UINTTAGS._serialized_start=1102
  _UINTTAGS._serialized_end=1288
  _STRINGTAGS._serialized_start=1291
  _STRINGTAGS._serialized_end=4054
  _BYTESTAGS._serialized_start=4057
  _BYTESTAGS._serialized_end=4264
  _BOOLTAGS._serialized_start=4266
  _BOOLTAGS._serialized_end=4300
  _ENUMTAGS._serialized_start=4303
  _ENUMTAGS._serialized_end=4525
  _MESSAGETAGS._serialized_start=4527
  _MESSAGETAGS._serialized_end=4604
  _REPEATEDTAGS._serialized_start=4607
  _REPEATEDTAGS._serialized_end=4922
  _MAPTAGS._serialized_start=4925
  _MAPTAGS._serialized_end=5240
# @@protoc_insertion_point(module_scope)
