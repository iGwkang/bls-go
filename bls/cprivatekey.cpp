extern "C" {
#include "cprivatekey.h"
}
#include "privatekey.hpp"

CPrivateKey PrivateKeyFromBytes(uint8_t bytes[PRIVATE_KEY_SIZE])
{
	return new bls::PrivateKey(bls::PrivateKey::FromBytes(bls::Bytes((uint8_t*)bytes, PRIVATE_KEY_SIZE)));
}
void FreePrivateKey(CPrivateKey pk)
{
    delete (bls::PrivateKey*)pk;
}

uint8_t * PrivateKeySerialize(CPrivateKey pk)
{
	uint8_t * mc = (uint8_t *)calloc(1, PRIVATE_KEY_SIZE);
	((bls::PrivateKey *)pk)->Serialize(mc);
	return mc;
}

CPrivateKey PrivateKeyCopy(CPrivateKey pk)
{
	return new bls::PrivateKey(*(bls::PrivateKey *)pk);
}

CG1Element PrivateKeyGetG1(CPrivateKey pk)
{
	return (CG1Element)&((bls::PrivateKey *)pk)->GetG1Element();
}

CG2Element PrivateKeyGetG2(CPrivateKey pk)
{
	return (CG2Element)&((bls::PrivateKey *)pk)->GetG2Element();
}

CPrivateKey PrivateKeyAggregate(CPrivateKey *vec, int count)
{
	std::vector<bls::PrivateKey> vecKey;
	for (int i = 0; i < count; ++i)
	{
		vecKey.emplace_back(*(bls::PrivateKey *)vec[i]);
	}
	return new bls::PrivateKey(bls::PrivateKey::Aggregate(vecKey));
}

bool EqualPrivateKey(CPrivateKey l, CPrivateKey r)
{
	return *(bls::PrivateKey *)l == *(bls::PrivateKey *)r;
}

CG1Element MulG1PrivateKey(CPrivateKey pk, CG1Element c1)
{
	return new bls::G1Element(*(bls::PrivateKey *)pk * *(bls::G1Element *)c1);
}

CG2Element MulG2PrivateKey(CPrivateKey pk, CG2Element c2)
{
	return new bls::G2Element(*(bls::PrivateKey *)pk * *(bls::G2Element *)c2);
}

