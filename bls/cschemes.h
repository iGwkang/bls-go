#pragma once
#include <stdint.h>

#include "relic_conf.h"
#include "celements.h"
#include "cprivatekey.h"

typedef struct CData
{
	uint8_t *data;
	int len;
} CData;

// BasicScheme
CG1Element BasicSchemeSkToG1(const CPrivateKey);
CPrivateKey BasicSchemeKeyGen(const uint8_t *data, int len);
CPrivateKey BasicSchemeDeriveChildSk(CPrivateKey, int index);
CPrivateKey BasicSchemeDeriveChildSkUnhardened(CPrivateKey, int index);
CG1Element BasicSchemeDeriveChildPkUnhardened(CG1Element, int index);
CG1Element BasicSchemeAggregateG1(CG1Element * vec, int count);
CG2Element BasicSchemeAggregateG2(CG2Element * vec, int count);
CG2Element BasicSchemeSign(CPrivateKey, const uint8_t *data, int len);
bool BasicSchemeVerify(CG1Element, const uint8_t *data, int len, CG2Element);
bool BasicSchemeeAggregateVerify(CG1Element *, int c1_count, const CData *data_vec, int len, CG2Element);

// AugScheme
CG1Element AugSchemeSkToG1(const CPrivateKey);
CPrivateKey AugSchemeKeyGen(const uint8_t *data, int len);
CPrivateKey AugSchemeDeriveChildSk(CPrivateKey, int index);
CPrivateKey AugSchemeDeriveChildSkUnhardened(CPrivateKey, int index);
CG1Element AugSchemeDeriveChildPkUnhardened(CG1Element, int index);
CG1Element AugSchemeAggregateG1(CG1Element * vec, int count);
CG2Element AugSchemeAggregateG2(CG2Element * vec, int count);
CG2Element AugSchemeSign(CPrivateKey, const uint8_t *data, int len);
bool AugSchemeVerify(CG1Element, const uint8_t *data, int len, CG2Element);
bool AugSchemeeAggregateVerify(CG1Element *, int c1_count, const CData *data_vec, int len, CG2Element);

// PopScheme
CG1Element PopSchemeSkToG1(const CPrivateKey);
CPrivateKey PopSchemeKeyGen(const uint8_t *data, int len);
CPrivateKey PopSchemeDeriveChildSk(CPrivateKey, int index);
CPrivateKey PopSchemeDeriveChildSkUnhardened(CPrivateKey, int index);
CG1Element PopSchemeDeriveChildPkUnhardened(CG1Element, int index);
CG1Element PopSchemeAggregateG1(CG1Element * vec, int count);
CG2Element PopSchemeAggregateG2(CG2Element * vec, int count);
CG2Element PopSchemeSign(CPrivateKey, const uint8_t *data, int len);
bool PopSchemeVerify(CG1Element, const uint8_t *data, int len, CG2Element);
bool PopSchemeeAggregateVerify(CG1Element *, int c1_count, const CData *data_vec, int len, CG2Element);
CG2Element PopSchemePopProve(CPrivateKey);
bool PopSchemePopVerify(CG1Element, CG2Element);
bool PopSchemeFastAggregateVerify(CG1Element *, int c1_count, const uint8_t *data, int len, CG2Element);



