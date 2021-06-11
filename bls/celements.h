#pragma once
#include <stdbool.h>

#ifdef __cplusplus
extern "C" {
#endif

#include "relic.h"
#include "relic_conf.h"
#include "util.h"


typedef void * CG1Element;

CG1Element NewG1Element();
void FreeG1Element(CG1Element);

CG1Element G1ElementCopy(CG1Element);

CG1Element G1ElementFromBytes(const uint8_t *data, int len);

CG1Element G1ElementFromMessage(const uint8_t *data, int len, const uint8_t *dst_data, int dst_len);

CG1Element G1ElementGenerator();

CG1Element G1ElementNegate(CG1Element);

uint32_t G1ElementGetFingerprint(CG1Element);

// 使用完需要调用者调用free
void G1ElementSerialize(CG1Element, uint8_t **out_data, uint32_t *out_size);

// ==
bool EqualG1Element(CG1Element lhs, CG1Element rhs);
// +
CG1Element AddG1Element(CG1Element self, CG1Element other);
// *
CG1Element MulG1Element(CG1Element self, const bn_t *);

// -------------------------------------------------------------------------------- //

typedef void * CG2Element;

CG2Element NewG2Element();
void FreeG2Element(CG2Element);

CG2Element G2ElementCopy(CG2Element);

CG2Element G2ElementFromBytes(const uint8_t *data, int len);

CG2Element G2ElementFromMessage(const uint8_t *data, int len, const uint8_t *dst_data, int dst_len);

CG2Element G2ElementGenerator();

CG2Element G2ElementNegate(CG2Element);

// 使用完需要调用者调用free
void G2ElementSerialize(CG2Element, uint8_t **out_data, uint32_t *out_size);

// ==
bool EqualG2Element(CG2Element lhs, CG2Element rhs);
// +
CG2Element AddG2Element(CG2Element self, CG2Element other);
// *
CG2Element MulG2Element(CG2Element self, const bn_t *);

#ifdef __cplusplus
}
#endif