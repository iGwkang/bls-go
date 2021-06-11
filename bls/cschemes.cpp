extern "C" {
#include "cschemes.h"
}
#include "schemes.hpp"

CG1Element BasicSchemeSkToG1(const CPrivateKey pk)
{
	return new bls::G1Element(bls::BasicSchemeMPL().SkToG1(*(bls::PrivateKey*)pk));
}

CPrivateKey BasicSchemeKeyGen(const uint8_t *data, int len)
{
	return new bls::PrivateKey(bls::BasicSchemeMPL().KeyGen(bls::Bytes(data, len)));
}

CPrivateKey BasicSchemeDeriveChildSk(CPrivateKey pk, int index)
{
	return new bls::PrivateKey(bls::BasicSchemeMPL().DeriveChildSk(*(bls::PrivateKey*)pk, index));
}

CPrivateKey BasicSchemeDeriveChildSkUnhardened(CPrivateKey pk, int index)
{
	return new bls::PrivateKey(bls::BasicSchemeMPL().DeriveChildSkUnhardened(*(bls::PrivateKey*)pk, index));
}

CG1Element BasicSchemeDeriveChildPkUnhardened(CG1Element c1, int index)
{
	return new bls::G1Element(bls::BasicSchemeMPL().DeriveChildPkUnhardened(*(bls::G1Element*)c1, index));
}

CG1Element BasicSchemeAggregateG1(CG1Element * vec, int count)
{
	std::vector<bls::G1Element> g1vec;
	for (int i = 0; i < count; ++i)
	{
		g1vec.emplace_back(*(bls::G1Element *)vec[i]);
	}

	return new bls::G1Element(bls::BasicSchemeMPL().Aggregate(g1vec));
}

CG2Element BasicSchemeAggregateG2(CG2Element * vec, int count)
{
	std::vector<bls::G2Element> g1vec;
	for (int i = 0; i < count; ++i)
	{
		g1vec.emplace_back(*(bls::G2Element *)vec[i]);
	}

	return new bls::G2Element(bls::BasicSchemeMPL().Aggregate(g1vec));
}

CG2Element BasicSchemeSign(CPrivateKey pk, const uint8_t *data, int len)
{
	return new bls::G2Element(bls::BasicSchemeMPL().Sign(*(bls::PrivateKey*)pk, bls::Bytes(data, len)));
}

bool BasicSchemeVerify(CG1Element c1, const uint8_t *data, int len, CG2Element c2)
{
	return bls::BasicSchemeMPL().Verify(*(bls::G1Element*)c1, bls::Bytes(data, len), *(bls::G2Element*)c2);
}

bool BasicSchemeeAggregateVerify(CG1Element *c1, int c1_count, const CData *data_vec, int len, CG2Element c2)
{
	std::vector<bls::G1Element> c1_vec;
	for (int i = 0; i < c1_count; ++i)
	{
		c1_vec.push_back(*(bls::G1Element*)c1[i]);
	}

	std::vector<bls::Bytes> data_v;
	for (int i = 0; i < len; ++i)
	{
		data_v.push_back(bls::Bytes(data_vec[i].data, data_vec[i].len));
	}

	return bls::BasicSchemeMPL().AggregateVerify(c1_vec, data_v, *(bls::G2Element*)c2);
}

CG1Element AugSchemeSkToG1(const CPrivateKey pk)
{
	return new bls::G1Element(bls::AugSchemeMPL().SkToG1(*(bls::PrivateKey*)pk));
}

CPrivateKey AugSchemeKeyGen(const uint8_t *data, int len)
{
	return new bls::PrivateKey(bls::AugSchemeMPL().KeyGen(bls::Bytes(data, len)));
}

CPrivateKey AugSchemeDeriveChildSk(CPrivateKey pk, int index)
{
	return new bls::PrivateKey(bls::AugSchemeMPL().DeriveChildSk(*(bls::PrivateKey*)pk, index));
}

CPrivateKey AugSchemeDeriveChildSkUnhardened(CPrivateKey pk, int index)
{
	return new bls::PrivateKey(bls::AugSchemeMPL().DeriveChildSkUnhardened(*(bls::PrivateKey*)pk, index));
}

CG1Element AugSchemeDeriveChildPkUnhardened(CG1Element c1, int index)
{
	return new bls::G1Element(bls::AugSchemeMPL().DeriveChildPkUnhardened(*(bls::G1Element*)c1, index));
}

CG1Element AugSchemeAggregateG1(CG1Element * vec, int count)
{
	std::vector<bls::G1Element> g1vec;
	for (int i = 0; i < count; ++i)
	{
		g1vec.emplace_back(*(bls::G1Element *)vec[i]);
	}

	return new bls::G1Element(bls::AugSchemeMPL().Aggregate(g1vec));
}

CG2Element AugSchemeAggregateG2(CG2Element * vec, int count)
{
	std::vector<bls::G2Element> g1vec;
	for (int i = 0; i < count; ++i)
	{
		g1vec.emplace_back(*(bls::G2Element *)vec[i]);
	}

	return new bls::G2Element(bls::AugSchemeMPL().Aggregate(g1vec));
}

CG2Element AugSchemeSign(CPrivateKey pk, const uint8_t *data, int len)
{
	return new bls::G2Element(bls::BasicSchemeMPL().Sign(*(bls::PrivateKey*)pk, bls::Bytes(data, len)));
}

bool AugSchemeVerify(CG1Element c1, const uint8_t *data, int len, CG2Element c2)
{
	return bls::AugSchemeMPL().Verify(*(bls::G1Element*)c1, bls::Bytes(data, len), *(bls::G2Element*)c2);
}

bool AugSchemeeAggregateVerify(CG1Element *c1, int c1_count, const CData *data_vec, int len, CG2Element c2)
{
	std::vector<bls::G1Element> c1_vec;
	for (int i = 0; i < c1_count; ++i)
	{
		c1_vec.push_back(*(bls::G1Element*)c1[i]);
	}

	std::vector<bls::Bytes> data_v;
	for (int i = 0; i < len; ++i)
	{
		data_v.push_back(bls::Bytes(data_vec[i].data, data_vec[i].len));
	}

	return bls::AugSchemeMPL().AggregateVerify(c1_vec, data_v, *(bls::G2Element*)c2);
}

CG1Element PopSchemeSkToG1(const CPrivateKey pk)
{
	return new bls::G1Element(bls::PopSchemeMPL().SkToG1(*(bls::PrivateKey*)pk));
}

CPrivateKey PopSchemeKeyGen(const uint8_t *data, int len)
{
	return new bls::PrivateKey(bls::PopSchemeMPL().KeyGen(bls::Bytes(data, len)));
}

CPrivateKey PopSchemeDeriveChildSk(CPrivateKey pk, int index)
{
	return new bls::PrivateKey(bls::PopSchemeMPL().DeriveChildSk(*(bls::PrivateKey*)pk, index));
}

CPrivateKey PopSchemeDeriveChildSkUnhardened(CPrivateKey pk, int index)
{
	return new bls::PrivateKey(bls::PopSchemeMPL().DeriveChildSkUnhardened(*(bls::PrivateKey*)pk, index));
}

CG1Element PopSchemeDeriveChildPkUnhardened(CG1Element c1, int index)
{
	return new bls::G1Element(bls::PopSchemeMPL().DeriveChildPkUnhardened(*(bls::G1Element*)c1, index));
}

CG1Element PopSchemeAggregateG1(CG1Element * vec, int count)
{
	std::vector<bls::G1Element> g1vec;
	for (int i = 0; i < count; ++i)
	{
		g1vec.emplace_back(*(bls::G1Element *)vec[i]);
	}

	return new bls::G1Element(bls::PopSchemeMPL().Aggregate(g1vec));
}

CG2Element PopSchemeAggregateG2(CG2Element * vec, int count)
{
	std::vector<bls::G2Element> g2vec;
	for (int i = 0; i < count; ++i)
	{
		g2vec.emplace_back(*(bls::G2Element *)vec[i]);
	}

	return new bls::G2Element(bls::PopSchemeMPL().Aggregate(g2vec));
}

CG2Element PopSchemeSign(CPrivateKey pk, const uint8_t *data, int len)
{
	return new bls::G2Element(bls::PopSchemeMPL().Sign(*(bls::PrivateKey*)pk, bls::Bytes(data, len)));
}

bool PopSchemeVerify(CG1Element c1, const uint8_t *data, int len, CG2Element c2)
{
	return bls::PopSchemeMPL().Verify(*(bls::G1Element*)c1, bls::Bytes(data, len), *(bls::G2Element*)c2);
}

bool PopSchemeeAggregateVerify(CG1Element * c1, int c1_count, const CData *data_vec, int len, CG2Element c2)
{
	std::vector<bls::G1Element> c1_vec;
	for (int i = 0; i < c1_count; ++i)
	{
		c1_vec.push_back(*(bls::G1Element*)c1[i]);
	}

	std::vector<bls::Bytes> data_v;
	for (int i = 0; i < len; ++i)
	{
		data_v.push_back(bls::Bytes(data_vec[i].data, data_vec[i].len));
	}

	return bls::PopSchemeMPL().AggregateVerify(c1_vec, data_v, *(bls::G2Element*)c2);
}

CG2Element PopSchemePopProve(CPrivateKey pk)
{
	return new bls::G2Element(bls::PopSchemeMPL().PopProve(*(bls::PrivateKey*)pk));
}

bool PopSchemePopVerify(CG1Element c1, CG2Element c2)
{
	return bls::PopSchemeMPL().PopVerify(*(bls::G1Element *)c1, *(bls::G2Element *)c2);
}

bool PopSchemeFastAggregateVerify(CG1Element * c1, int c1_count, const uint8_t *data, int len, CG2Element c2)
{
	std::vector<bls::G1Element> c1_vec;
	for (int i = 0; i < c1_count; ++i)
	{
		c1_vec.push_back(*(bls::G1Element*)c1[i]);
	}
	return bls::PopSchemeMPL().FastAggregateVerify(c1_vec, bls::Bytes(data, len), *(bls::G2Element *)c2);
}

