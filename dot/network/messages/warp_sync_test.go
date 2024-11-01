// Copyright 2024 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package messages

import (
	"encoding/hex"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecodeWarpSyncProof(t *testing.T) {
	// Generated using substrate
	hexString := "04e995dd4dddbdfc54bd8126e0b5b394a47da71ca4a6510e7bd132d103470bf9b9be748b05d4026dd4bd8821893ce9e82a3a8c324ca454dd7510ee9ab05df3df82e87a1eada450bc10f223a562a12962e8914f436fd0991319d662310298fa47e5ff0dffb50c0642414245b501031100000009ba301100000000a6815e866aebee63d23cefa60964fa88b4814264fd35703c0c8a5f8192d8f41aa15bc5d991d1cad65ffccb73dbacd542a32286ec051190166ec00bfb4304b6049a882d0858cdbfb8d37a07269ea01bad1b55fb1174d6dedbfac56787adfa700004424545468403a6c0199af958a3ae8a3928b0802b709589bcbb1ec5b7041cf91644dd696ad99f05424142450101f8b400261b84cc580a8947abe4caf3b4ee51b7aeb26afc9a3bcdd2b3cd634f2df03cda49368e5e56a5b2427f6088cdc4144d75e1985f5e0a7e8f527290f47b8dad10000000000000376d697522a9a42b9c6952e771fd74be9cb19fe6dd0d41c362f753ae2686547a2fdd620138376d697522a9a42b9c6952e771fd74be9cb19fe6dd0d41c362f753ae2686547a2fdd62015037bddf7847fd40d7e6235c17203c8b70e661ef25aa44cbbaed19d5bb455bec3b20f59a720253b7c536171e1e68a91dcbf6b88fd077e28a2125cc9ea4fb430819b6ba718c28b1078474dc63e9d8a6f4d794da79f661da070fcca13b3bba50e7376d697522a9a42b9c6952e771fd74be9cb19fe6dd0d41c362f753ae2686547a2fdd6201243478b0465e1c13e9cda8479d42958422ab0acd69a1dbdc8d14fa95d2b15e96e6c1a5e58df7ff148977ba68418a7d6f6cf2ab959056c59cd768e490bc90ea0c1a619f90a9c96ec3e1c57211613247e8c0efe123188263da1dc8796d6a018c16376d697522a9a42b9c6952e771fd74be9cb19fe6dd0d41c362f753ae2686547a2fdd620151010e386ecf4cb60e88acf75d7859547584a8e4258af8ca3fcde8d79dedf9ed1b6b0bf7df2704c3dfbc2db4a785833f0a70e30ec3c6ad9f35cae93e6ebd990a1d50350ac619dc36a0fe05cdd7a1654e73858586d55265f3524cc27e005be0f2376d697522a9a42b9c6952e771fd74be9cb19fe6dd0d41c362f753ae2686547a2fdd6201ca80d95b84408c14cc2c0cca1f3d2e094ebe056b5816959c887ac40b7311edf8c87e0f77a06ed398323b684e09c6d784e5333136dbae5a1d0627bf774e633408229366210b4b20adf7741fb6bc8fa669ea2f7470e1440b38969d2a47d752ffeb376d697522a9a42b9c6952e771fd74be9cb19fe6dd0d41c362f753ae2686547a2fdd6201bda1ac4056f6e6210d6b2dd8f69323cb7f2729f49ad9d42059ceecdffecebef04ce478a3cfde9bceadaf9be978443f0c3a5aea19d04b538b087cc0ca614e6b0a2bf793d9fb4251da094a2dd33bb0f68a46e60d799b58a78a4a5cc369a92ee2ec376d697522a9a42b9c6952e771fd74be9cb19fe6dd0d41c362f753ae2686547a2fdd620156dd1808cb6044b9d6762bbd5a0284b32e51809010300fa8b4a6f2bb3072e9e632041c6ad74ba989c3e18f3874ab4be95c2e8f51ab275d943a839c12d5b73c0a3ab4559bf186773ede8cf3ddc78a93439c8e82ead090de7d3c2fe4e75591fcf1376d697522a9a42b9c6952e771fd74be9cb19fe6dd0d41c362f753ae2686547a2fdd620176857f28231e3b3a8eaf85eb07b835f5f7f5d379e854debd3f325693118db758b2765fa1c4a959d41f7d0b1650f61cffe1d607eee7fb12de81777cf725cd8e073fee6e433392b821f11c4cf8fe8d17b7c08dbbc32381895955ee27448f8869e1376d697522a9a42b9c6952e771fd74be9cb19fe6dd0d41c362f753ae2686547a2fdd6201688b6d2c8640491fe4db52979dfaa2947ebded2b4ba5e020257fe29ffdb2a6c7017102fa6c9d6f0744bba28a05c34c377108bf73ee04982d6db0656208ec620c7e287589ac74a46b3a74a8ae66c3e2554857419bd00f9bfeb55957fffccd98b9376d697522a9a42b9c6952e771fd74be9cb19fe6dd0d41c362f753ae2686547a2fdd62012256e07d9c9f3c76356fa7946b8233591a4d9b91ed3e81b2a09b1d3ae6b69c8906a299e8245221c8efd46bc55966f1de6b2b4bd0b2b3256527d59e7b738056009e8adb538dacdd20c86c4eca9a158895cef125dfd118e883cd55454f8d59f3ed376d697522a9a42b9c6952e771fd74be9cb19fe6dd0d41c362f753ae2686547a2fdd62011479e687ce7b0594637250707491cbec69a909aad5dab78e00aae8c3d8a530d8de4c6a1faad6ef416aea872988d50227e35e2dfb5b73d4b3b897acaa4cf56f0ea674568468af9fef031f033166880eec9f5ddd05d797d5abf1a9b9a957c77820376d697522a9a42b9c6952e771fd74be9cb19fe6dd0d41c362f753ae2686547a2fdd6201a0e495db4830b2a4e80d045871b5f0f8541091fdb7707d5e6c8f085bc255f0a5069a0ae153d9a83f6fa3f208cd1d701fbd7e852449f026f8566c7d8826cf5b0aa99c0755eec29f2be753fb701762d1e7cc841323bac49576b9ea2e124c4b7b9f376d697522a9a42b9c6952e771fd74be9cb19fe6dd0d41c362f753ae2686547a2fdd62018c4f486b7b87c5735e42249c15f5ec6d7b2b50a16aa0ba3bfbe666f99537e81a6a121659a06d58eacb9594201b6644f18e54c1c6ca457a4fcfa9d15e3bfea201be0bbf54aad305682afc6dc7cd051b5dce154e7d0056e877562407872cdb1242376d697522a9a42b9c6952e771fd74be9cb19fe6dd0d41c362f753ae2686547a2fdd6201763fba0da4fef393f59f1fd9eb9a646a2a1106f404681840fa846ec56c1a3443223b8a147990c6fd3e2af11300531168be94f091931f6ac07724f3ce75a06600d3ccc285f3c648f34ca9ce476651495b9a6c4aa7607d83737bd871de56a50ed9376d697522a9a42b9c6952e771fd74be9cb19fe6dd0d41c362f753ae2686547a2fdd6201d9b7c0218e498e5930fc4ad9a8c139f774de3fda36ad408008c3957f1e822070ceda5ca810ef3fb7fda2c48576051107235edba92a3be39b38b9ac9aba8c2305f12e7ddae48873b7dcd614b8ccc8ea2e5ec000f3cc398f20801cf9047bc909b20001"
	expected, err := hex.DecodeString(hexString)
	if err != nil {
		log.Fatal(err)
	}

	var proof WarpSyncProof

	err = proof.Decode(expected)
	require.NoError(t, err)

	encoded, err := proof.Encode()
	require.NoError(t, err)
	require.Equal(t, expected, encoded)
}
