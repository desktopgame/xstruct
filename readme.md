# xstruct
xstructは、xmlファイルをもとに対応するメタデータ付きの構造体と
そのマーシャリング/アンマーシャリング関数を生成するツールです。

# install
````
go get github.com/desktopgame/xstruct
````

# how to use
````
xstruct data.xml
xstruct -package mypkg -prefix=Prefix -suffix=Suffix data.xml
````

# sample
VisualStudioのnuget定義ファイルを構造体へ変換する

**INPUT**
````
<?xml version="1.0" encoding="utf-8"?>
<packages>
  <package id="freeglut" version="2.8.1.15" targetFramework="native" />
  <package id="freeglut.redist" version="2.8.1.15" targetFramework="native" />
  <package id="freetype" version="2.8.0.1" targetFramework="native" />
  <package id="freetype.redist" version="2.8.0.1" targetFramework="native" />
  <package id="glew" version="1.9.0.1" targetFramework="native" />
  <package id="glew.redist" version="1.9.0.1" targetFramework="native" />
  <package id="glfw" version="3.3.0.1" targetFramework="native" />
  <package id="glm" version="0.9.9.600" targetFramework="native" />
  <package id="soil" version="1.16.0" targetFramework="native" />
</packages>
````
**OUTPUT**
````
package main
type Packages struct {
    // define attribute
    // define subelement
    SubPackage []*PackagesPackage `xml:"package"`
    // define content
    Content string `xml:",chardata"`
}
type PackagesPackage struct {
    // define attribute
    AttrtargetFramework string `xml:"targetFramework,attr"`
    Attrid string `xml:"id,attr"`
    Attrversion string `xml:"version,attr"`
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
func LoadPackages(path string) (*Packages, error) {
    xmlFile, err := os.Open(path)
    if err != nil {
    	return nil, err
    }
    defer xmlFile.Close()
    xmlData, err := ioutil.ReadAll(xmlFile)
    if err != nil {
    	return nil, err
    }
    var data Packages
    xml.Unmarshal(xmlData, &data)
    return &data, nil
}

func SavePackages(path string, data *Packages, perm os.FileMode) error {
    buf, err := xml.MarshalIndent(data, "", "    ")
    if err != nil {
        return err
    }
    err = ioutil.WriteFile(path, buf, perm)
    if err != nil {
    	return err
    }
    return nil
}
````

# sample2
もう少し大きなサンプル
VisualStudioのプロジェクトを構造体へ変換する

**INPUT**
````
<?xml version="1.0" encoding="utf-8"?>
<Project DefaultTargets="Build" ToolsVersion="15.0" xmlns="http://schemas.microsoft.com/developer/msbuild/2003">
  <Import Project="..\packages\soil.1.16.0\build\native\soil.props" Condition="Exists('..\packages\soil.1.16.0\build\native\soil.props')" />
  <ItemGroup Label="ProjectConfigurations">
    <ProjectConfiguration Include="Debug|Win32">
      <Configuration>Debug</Configuration>
      <Platform>Win32</Platform>
    </ProjectConfiguration>
    <ProjectConfiguration Include="Release|Win32">
      <Configuration>Release</Configuration>
      <Platform>Win32</Platform>
    </ProjectConfiguration>
    <ProjectConfiguration Include="Debug|x64">
      <Configuration>Debug</Configuration>
      <Platform>x64</Platform>
    </ProjectConfiguration>
    <ProjectConfiguration Include="Release|x64">
      <Configuration>Release</Configuration>
      <Platform>x64</Platform>
    </ProjectConfiguration>
  </ItemGroup>
  <PropertyGroup Label="Globals">
    <VCProjectVersion>15.0</VCProjectVersion>
    <ProjectGuid>{5A11DAC0-4F35-4BF8-AB39-AD2489531C1C}</ProjectGuid>
    <RootNamespace>Planet</RootNamespace>
    <WindowsTargetPlatformVersion>10.0.17763.0</WindowsTargetPlatformVersion>
  </PropertyGroup>
  <Import Project="$(VCTargetsPath)\Microsoft.Cpp.Default.props" />
  <PropertyGroup Condition="'$(Configuration)|$(Platform)'=='Debug|Win32'" Label="Configuration">
    <ConfigurationType>Application</ConfigurationType>
    <UseDebugLibraries>true</UseDebugLibraries>
    <PlatformToolset>v141</PlatformToolset>
    <CharacterSet>MultiByte</CharacterSet>
  </PropertyGroup>
  <PropertyGroup Condition="'$(Configuration)|$(Platform)'=='Release|Win32'" Label="Configuration">
    <ConfigurationType>Application</ConfigurationType>
    <UseDebugLibraries>false</UseDebugLibraries>
    <PlatformToolset>v141</PlatformToolset>
    <WholeProgramOptimization>true</WholeProgramOptimization>
    <CharacterSet>MultiByte</CharacterSet>
  </PropertyGroup>
  <PropertyGroup Condition="'$(Configuration)|$(Platform)'=='Debug|x64'" Label="Configuration">
    <ConfigurationType>Application</ConfigurationType>
    <UseDebugLibraries>true</UseDebugLibraries>
    <PlatformToolset>v141</PlatformToolset>
    <CharacterSet>MultiByte</CharacterSet>
  </PropertyGroup>
  <PropertyGroup Condition="'$(Configuration)|$(Platform)'=='Release|x64'" Label="Configuration">
    <ConfigurationType>Application</ConfigurationType>
    <UseDebugLibraries>false</UseDebugLibraries>
    <PlatformToolset>v141</PlatformToolset>
    <WholeProgramOptimization>true</WholeProgramOptimization>
    <CharacterSet>MultiByte</CharacterSet>
  </PropertyGroup>
  <Import Project="$(VCTargetsPath)\Microsoft.Cpp.props" />
  <ImportGroup Label="ExtensionSettings">
  </ImportGroup>
  <ImportGroup Label="Shared">
  </ImportGroup>
  <ImportGroup Label="PropertySheets" Condition="'$(Configuration)|$(Platform)'=='Debug|Win32'">
    <Import Project="$(UserRootDir)\Microsoft.Cpp.$(Platform).user.props" Condition="exists('$(UserRootDir)\Microsoft.Cpp.$(Platform).user.props')" Label="LocalAppDataPlatform" />
  </ImportGroup>
  <ImportGroup Label="PropertySheets" Condition="'$(Configuration)|$(Platform)'=='Release|Win32'">
    <Import Project="$(UserRootDir)\Microsoft.Cpp.$(Platform).user.props" Condition="exists('$(UserRootDir)\Microsoft.Cpp.$(Platform).user.props')" Label="LocalAppDataPlatform" />
  </ImportGroup>
  <ImportGroup Label="PropertySheets" Condition="'$(Configuration)|$(Platform)'=='Debug|x64'">
    <Import Project="$(UserRootDir)\Microsoft.Cpp.$(Platform).user.props" Condition="exists('$(UserRootDir)\Microsoft.Cpp.$(Platform).user.props')" Label="LocalAppDataPlatform" />
  </ImportGroup>
  <ImportGroup Label="PropertySheets" Condition="'$(Configuration)|$(Platform)'=='Release|x64'">
    <Import Project="$(UserRootDir)\Microsoft.Cpp.$(Platform).user.props" Condition="exists('$(UserRootDir)\Microsoft.Cpp.$(Platform).user.props')" Label="LocalAppDataPlatform" />
  </ImportGroup>
  <PropertyGroup Label="UserMacros" />
  <PropertyGroup />
  <ItemDefinitionGroup Condition="'$(Configuration)|$(Platform)'=='Debug|Win32'">
    <ClCompile>
      <WarningLevel>Level3</WarningLevel>
      <Optimization>Disabled</Optimization>
      <SDLCheck>true</SDLCheck>
      <ConformanceMode>true</ConformanceMode>
      <AdditionalIncludeDirectories>$(SolutionDir)packages\freetype.2.8.0.1\build\native\include;$(SolutionDir)packages\soil.1.16.0\build\native\include;$(SolutionDir)packages\glm.0.9.9.500\build\native\include;C:\Program Files\Autodesk\FBX\FBX SDK\2019.2\include;C:\Program Files (x86)\OpenAL 1.1 SDK\ALUT\freealut-freealut_1_1_0\include;C:\Program Files (x86)\OpenAL 1.1 SDK\include;$(SolutionDir)packages\freeglut.2.8.1.15\build\native\include\GL;$(SolutionDir)packages\glew.1.9.0.1\build\native\include\GL;$(SolutionDir)packages\glfw.3.3.0.1\build\native\include\GLFW;%(AdditionalIncludeDirectories)</AdditionalIncludeDirectories>
      <LanguageStandard>stdcpp17</LanguageStandard>
      <PreprocessorDefinitions>_CRT_SECURE_NO_WARNINGS;DEBUG;_MBCS;%(PreprocessorDefinitions)</PreprocessorDefinitions>
      <RuntimeLibrary>MultiThreadedDebugDLL</RuntimeLibrary>
    </ClCompile>
    <Link>
      <AdditionalLibraryDirectories>$(SolutionDir)packages\freetype.2.8.0.1\build\native\lib\Win32\v141\static\Debug;$(SolutionDir)packages\soil.1.16.0\build\native\lib\Win32\Debug;C:\Program Files\Autodesk\FBX\FBX SDK\2019.2\lib\vs2017\x86\debug;C:\Program Files (x86)\OpenAL 1.1 SDK\ALUT\build\debug;C:\Program Files (x86)\OpenAL 1.1 SDK\libs\Win32\EFX-Util_MT;C:\Program Files (x86)\OpenAL 1.1 SDK\libs\Win32;$(SolutionDir)packages\zlib.v140.windesktop.msvcstl.dyn.rt-dyn.1.2.8.8\lib\native\v140\windesktop\msvcstl\dyn\rt-dyn\Win32\Debug;$(SolutionDir)packages\glfw.3.3.0.1\build\native\lib\static\v142\win32;$(SolutionDir)packages\glew.1.9.0.1\build\native\lib\v110\Win32\Debug\static;$(SolutionDir)packages\freeglut.2.8.1.15\build\native\lib\v110\Win32\Debug\static;%(AdditionalLibraryDirectories)</AdditionalLibraryDirectories>
      <AdditionalDependencies>freetype28d.lib;freeglut.lib;glew.lib;glfw3.lib;zlib-md.lib;alut.lib;OpenAL32.lib;EFX-Util.lib;libfbxsdk.lib;libxml2-md.lib;soil.lib;%(AdditionalDependencies)</AdditionalDependencies>
    </Link>
    <PreBuildEvent>
      <Command>
      </Command>
    </PreBuildEvent>
  </ItemDefinitionGroup>
  <ItemDefinitionGroup Condition="'$(Configuration)|$(Platform)'=='Debug|x64'">
    <ClCompile>
      <WarningLevel>Level3</WarningLevel>
      <Optimization>Disabled</Optimization>
      <SDLCheck>true</SDLCheck>
      <ConformanceMode>true</ConformanceMode>
      <AdditionalIncludeDirectories>$(SolutionDir)packages\libpng.1.6.28.1\build\native\include;$(SolutionDir)packages\glfw.3.3.0.1\build\native\include\GLFW;$(SolutionDir)packages\glew.1.9.0.1\build\native\include\GL;$(SolutionDir)packages\freeglut.2.8.1.15\build\native\include\GL;C:\Program Files (x86)\OpenAL 1.1 SDK\include;C:\Program Files (x86)\OpenAL 1.1 SDK\ALUT\freealut-freealut_1_1_0\include;C:\Program Files\Autodesk\FBX\FBX SDK\2019.2\include;%(AdditionalIncludeDirectories)</AdditionalIncludeDirectories>
    </ClCompile>
    <Link>
      <AdditionalLibraryDirectories>$(SolutionDir)packages\freeglut.2.8.1.15\build\native\lib\v110\Win32\Debug\static;$(SolutionDir)packages\glew.1.9.0.1\build\native\lib\v110\Win32\Debug\static;$(SolutionDir)packages\glfw.3.3.0.1\build\native\lib\static\v142\win32;$(SolutionDir)packages\libpng.1.6.28.1\build\native\lib\Win32\v140\dynamic\Debug;$(SolutionDir)packages\zlib.v140.windesktop.msvcstl.dyn.rt-dyn.1.2.8.8\lib\native\v140\windesktop\msvcstl\dyn\rt-dyn\Win32\Debug;C:\Program Files (x86)\OpenAL 1.1 SDK\libs\Win32;C:\Program Files (x86)\OpenAL 1.1 SDK\libs\Win32\EFX-Util_MT;C:\Program Files (x86)\OpenAL 1.1 SDK\ALUT\build\debug;C:\Program Files\Autodesk\FBX\FBX SDK\2019.2\lib\vs2017\x86\debug;%(AdditionalLibraryDirectories)</AdditionalLibraryDirectories>
      <AdditionalDependencies>freeglut.lib;glew.lib;glfw3.lib;zlibd.lib;libpng16.lib;%(AdditionalDependencies)</AdditionalDependencies>
    </Link>
  </ItemDefinitionGroup>
  <ItemDefinitionGroup Condition="'$(Configuration)|$(Platform)'=='Release|Win32'">
    <ClCompile>
      <WarningLevel>Level3</WarningLevel>
      <Optimization>MaxSpeed</Optimization>
      <FunctionLevelLinking>true</FunctionLevelLinking>
      <IntrinsicFunctions>true</IntrinsicFunctions>
      <SDLCheck>true</SDLCheck>
      <ConformanceMode>true</ConformanceMode>
      <AdditionalIncludeDirectories>$(SolutionDir)packages\libjpeg.9.2.0.1\build\native\include;C:\Work\PlanetSln\Planet\packages\glm.0.9.9.500\build\native\include;C:\Program Files\Autodesk\FBX\FBX SDK\2019.2\include;C:\Program Files (x86)\OpenAL 1.1 SDK\ALUT\freealut-freealut_1_1_0\include;C:\Program Files (x86)\OpenAL 1.1 SDK\include;$(SolutionDir)packages\freeglut.2.8.1.15\build\native\include\GL;$(SolutionDir)packages\glew.1.9.0.1\build\native\include\GL;$(SolutionDir)packages\glfw.3.3.0.1\build\native\include\GLFW;$(SolutionDir)packages\libpng.1.6.28.1\build\native\include;%(AdditionalIncludeDirectories)</AdditionalIncludeDirectories>
      <LanguageStandard>stdcpp17</LanguageStandard>
      <PreprocessorDefinitions>_CRT_SECURE_NO_WARNINGS;DEBUG;_MBCS;%(PreprocessorDefinitions)</PreprocessorDefinitions>
    </ClCompile>
    <Link>
      <EnableCOMDATFolding>true</EnableCOMDATFolding>
      <OptimizeReferences>true</OptimizeReferences>
      <AdditionalLibraryDirectories>$(SolutionDir)packages\libjpeg.9.2.0.1\build\native\lib\v140\Win32\Release\static\cdecl;C:\Program Files\Autodesk\FBX\FBX SDK\2019.2\lib\vs2017\x86\Release;C:\Program Files (x86)\OpenAL 1.1 SDK\ALUT\build\Release;C:\Program Files (x86)\OpenAL 1.1 SDK\libs\Win32\EFX-Util_MT;C:\Program Files (x86)\OpenAL 1.1 SDK\libs\Win32;$(SolutionDir)packages\zlib.v140.windesktop.msvcstl.dyn.rt-dyn.1.2.8.8\lib\native\v140\windesktop\msvcstl\dyn\rt-dyn\Win32\Release;$(SolutionDir)packages\libpng.1.6.28.1\build\native\lib\Win32\v140\dynamic\Release;$(SolutionDir)packages\glfw.3.3.0.1\build\native\lib\static\v142\win32;$(SolutionDir)packages\glew.1.9.0.1\build\native\lib\v110\Win32\Release\static;$(SolutionDir)packages\freeglut.2.8.1.15\build\native\lib\v110\Win32\Release\static;%(AdditionalLibraryDirectories)</AdditionalLibraryDirectories>
      <AdditionalDependencies>jpeg.lib;freeglut.lib;glew.lib;glfw3.lib;zlib-md.lib;libpng16.lib;alut.lib;OpenAL32.lib;EFX-Util.lib;libfbxsdk.lib;libxml2-md.lib;%(AdditionalDependencies)</AdditionalDependencies>
    </Link>
  </ItemDefinitionGroup>
  <ItemDefinitionGroup Condition="'$(Configuration)|$(Platform)'=='Release|x64'">
    <ClCompile>
      <WarningLevel>Level3</WarningLevel>
      <Optimization>MaxSpeed</Optimization>
      <FunctionLevelLinking>true</FunctionLevelLinking>
      <IntrinsicFunctions>true</IntrinsicFunctions>
      <SDLCheck>true</SDLCheck>
      <ConformanceMode>true</ConformanceMode>
    </ClCompile>
    <Link>
      <EnableCOMDATFolding>true</EnableCOMDATFolding>
      <OptimizeReferences>true</OptimizeReferences>
    </Link>
  </ItemDefinitionGroup>
  <ItemGroup>
    <ClCompile Include="game\input\KeyMove.cpp" />
    <ClCompile Include="game\input\MouseScroll.cpp" />
    <ClCompile Include="game\main.cpp" />
    <ClCompile Include="game\MyGame.cpp" />
    <ClCompile Include="game\resources.cpp" />
    <ClCompile Include="game\scene\LoadScene.cpp" />
    <ClCompile Include="game\scene\PlayScene.cpp" />
    <ClCompile Include="game\scene\TestScene.cpp" />
    <ClCompile Include="game\scene\TitleScene.cpp" />
    <ClCompile Include="game\scene\TutorialScene.cpp" />
    <ClCompile Include="game\ui\RightHandUI.cpp" />
    <ClCompile Include="game\world\biome\BasicBiome.cpp" />
    <ClCompile Include="game\world\biome\Biome.cpp" />
    <ClCompile Include="game\world\biome\DesertBiome.cpp" />
    <ClCompile Include="game\world\biome\HillBiome.cpp" />
    <ClCompile Include="game\world\biome\PlainBiome.cpp" />
    <ClCompile Include="game\world\Block.cpp" />
    <ClCompile Include="game\world\BlockRegistry.cpp" />
    <ClCompile Include="game\world\Entity.cpp" />
    <ClCompile Include="game\world\EntityPhysics.cpp" />
    <ClCompile Include="game\world\gen\Cell.cpp" />
    <ClCompile Include="game\world\gen\Xorshift.cpp" />
    <ClCompile Include="game\world\MultiBlock.cpp" />
    <ClCompile Include="game\world\Planet.cpp" />
    <ClCompile Include="game\world\Space.cpp" />
    <ClCompile Include="game\world\TexturePack.cpp" />
    <ClCompile Include="game\world\WarpPoint.cpp" />
    <ClCompile Include="game\world\World.cpp" />
    <ClCompile Include="gel\asset\AssetLoadException.cpp" />
    <ClCompile Include="gel\asset\FontLayout.cpp" />
    <ClCompile Include="gel\asset\FreeTypeFont.cpp" />
    <ClCompile Include="gel\asset\FreeTypeFontInstance.cpp" />
    <ClCompile Include="gel\asset\ProxyTexture.cpp" />
    <ClCompile Include="gel\GameConfig.cpp" />
    <ClCompile Include="gel\input\Drag.cpp" />
    <ClCompile Include="gel\input\Input.cpp" />
    <ClCompile Include="gel\input\InputState.cpp" />
    <ClCompile Include="gel\input\InputSystem.cpp" />
    <ClCompile Include="gel\input\KeyboardBuffer.cpp" />
    <ClCompile Include="gel\input\KeyboardState.cpp" />
    <ClCompile Include="gel\input\MouseState.cpp" />
    <ClCompile Include="gel\pipeline\BmpPipeline.cpp" />
    <ClCompile Include="gel\pipeline\ContentManager.cpp" />
    <ClCompile Include="gel\pipeline\FbxPipeline.cpp" />
    <ClCompile Include="gel\pipeline\JpegPipeline.cpp" />
    <ClCompile Include="gel\pipeline\OtfPipeline.cpp" />
    <ClCompile Include="gel\pipeline\PngPipeline.cpp" />
    <ClCompile Include="gel\pipeline\TtfPipeline.cpp" />
    <ClCompile Include="gel\pipeline\WavePipeline.cpp" />
    <ClCompile Include="gel\asset\AssetDatabase.cpp" />
    <ClCompile Include="gel\asset\FbxModel.cpp" />
    <ClCompile Include="gel\asset\SOILTexture.cpp" />
    <ClCompile Include="gel\asset\TextureIO.cpp" />
    <ClCompile Include="gel\asset\WaveAudio.cpp" />
    <ClCompile Include="gel\Game.cpp" />
    <ClCompile Include="gel\gsystem\Camera.cpp" />
    <ClCompile Include="gel\gsystem\Duration.cpp" />
    <ClCompile Include="gel\gsystem\LineRenderer.cpp" />
    <ClCompile Include="gel\math\Random.cpp" />
    <ClCompile Include="gel\gsystem\Timer.cpp" />
    <ClCompile Include="gel\gsystem\Transform.cpp" />
    <ClCompile Include="gel\math\AABB.cpp" />
    <ClCompile Include="gel\math\Quadrangle.cpp" />
    <ClCompile Include="gel\math\Rectangle.cpp" />
    <ClCompile Include="gel\math\Triangle.cpp" />
    <ClCompile Include="gel\scene\AnimationComponent.cpp" />
    <ClCompile Include="gel\scene\Component.cpp" />
    <ClCompile Include="gel\scene\GameObject.cpp" />
    <ClCompile Include="gel\scene\RendererComponent.cpp" />
    <ClCompile Include="gel\scene\Scene.cpp" />
    <ClCompile Include="gel\scene\SceneManager.cpp" />
    <ClCompile Include="gel\scene\SpriteComponent.cpp" />
    <ClCompile Include="gel\scene\TimerComponent.cpp" />
    <ClCompile Include="gel\scene\UIArrowButtonComponent.cpp" />
    <ClCompile Include="gel\scene\UIButtonComponent.cpp" />
    <ClCompile Include="gel\scene\UIElementComponent.cpp" />
    <ClCompile Include="gel\scene\UIFrameComponent.cpp" />
    <ClCompile Include="gel\scene\UIImageComponent.cpp" />
    <ClCompile Include="gel\scene\UILayoutComponent.cpp" />
    <ClCompile Include="gel\scene\UIProgressBarComponent.cpp" />
    <ClCompile Include="gel\scene\UIRect.cpp" />
    <ClCompile Include="gel\scene\UIRelativeLayoutComponent.cpp" />
    <ClCompile Include="gel\scene\UITextComponent.cpp" />
    <ClCompile Include="gel\scene\UITheme.cpp" />
    <ClCompile Include="gel\service\FBXSDKService.cpp" />
    <ClCompile Include="gel\service\FreeTypeService.cpp" />
    <ClCompile Include="gel\shader\Bitmap.cpp" />
    <ClCompile Include="gel\shader\Box.cpp" />
    <ClCompile Include="gel\shader\Circle.cpp" />
    <ClCompile Include="gel\shader\Color4.cpp" />
    <ClCompile Include="gel\shader\CubeMap.cpp" />
    <ClCompile Include="gel\shader\FontRenderer.cpp" />
    <ClCompile Include="gel\shader\FrameBuffer.cpp" />
    <ClCompile Include="gel\shader\IRMaterial.cpp" />
    <ClCompile Include="gel\shader\IRMesh.cpp" />
    <ClCompile Include="gel\shader\IRModel.cpp" />
    <ClCompile Include="gel\shader\IRShape.cpp" />
    <ClCompile Include="gel\shader\Layer.cpp" />
    <ClCompile Include="gel\shader\Line.cpp" />
    <ClCompile Include="gel\shader\NameRule.cpp" />
    <ClCompile Include="gel\shader\PixelBuffer.cpp" />
    <ClCompile Include="gel\shader\Plane.cpp" />
    <ClCompile Include="gel\shader\PlaneBatch.cpp" />
    <ClCompile Include="gel\shader\PlaneDraw.cpp" />
    <ClCompile Include="gel\shader\Primitive.cpp" />
    <ClCompile Include="gel\shader\RawTexture.cpp" />
    <ClCompile Include="gel\shader\Reflection4.cpp" />
    <ClCompile Include="gel\shader\RenderBuffer.cpp" />
    <ClCompile Include="gel\shader\Renderer.cpp" />
    <ClCompile Include="gel\shader\ScreenBuffer.cpp" />
    <ClCompile Include="gel\shader\Shader.cpp" />
    <ClCompile Include="gel\shader\ShaderRegistry.cpp" />
    <ClCompile Include="gel\shader\VertexArray.cpp" />
    <ClCompile Include="gel\signal\Connection.cpp" />
    <ClCompile Include="gel\util\CommandReader.cpp" />
    <ClCompile Include="gel\util\error.cpp" />
    <ClCompile Include="gel\util\Flag.cpp" />
    <ClCompile Include="gel\util\io.cpp" />
    <ClCompile Include="gel\util\KeyTrigger.cpp" />
    <ClCompile Include="gel\util\MouseTrigger.cpp" />
    <ClCompile Include="gel\util\Names.cpp" />
    <ClCompile Include="gel\util\string.cpp" />
    <ClCompile Include="gel\Window.cpp" />
  </ItemGroup>
  <ItemGroup>
    <ClInclude Include="game\def.hpp" />
    <ClInclude Include="game\gel.hpp" />
    <ClInclude Include="game\input\KeyMove.hpp" />
    <ClInclude Include="game\input\MouseScroll.hpp" />
    <ClInclude Include="game\MyGame.hpp" />
    <ClInclude Include="game\resources.hpp" />
    <ClInclude Include="game\scene\LoadScene.hpp" />
    <ClInclude Include="game\scene\PlayScene.hpp" />
    <ClInclude Include="game\scene\TestScene.hpp" />
    <ClInclude Include="game\scene\TitleScene.hpp" />
    <ClInclude Include="game\scene\TutorialScene.hpp" />
    <ClInclude Include="game\ui\RightHandUI.hpp" />
    <ClInclude Include="game\world\biome\BasicBiome.hpp" />
    <ClInclude Include="game\world\biome\Biome.hpp" />
    <ClInclude Include="game\world\biome\DesertBiome.hpp" />
    <ClInclude Include="game\world\biome\HillBiome.hpp" />
    <ClInclude Include="game\world\biome\PlainBiome.hpp" />
    <ClInclude Include="game\world\Block.hpp" />
    <ClInclude Include="game\world\BlockRegistry.hpp" />
    <ClInclude Include="game\world\Entity.hpp" />
    <ClInclude Include="game\world\EntityPhysics.hpp" />
    <ClInclude Include="game\world\gen\Cell.hpp" />
    <ClInclude Include="game\world\gen\Generator.hpp" />
    <ClInclude Include="game\world\gen\PerlinNoise.hpp" />
    <ClInclude Include="game\world\gen\Xorshift.hpp" />
    <ClInclude Include="game\world\MultiBlock.hpp" />
    <ClInclude Include="game\world\Planet.hpp" />
    <ClInclude Include="game\world\Space.hpp" />
    <ClInclude Include="game\world\TexturePack.hpp" />
    <ClInclude Include="game\world\WarpPoint.hpp" />
    <ClInclude Include="game\world\World.hpp" />
    <ClInclude Include="gel\all.hpp" />
    <ClInclude Include="gel\asset\AssetLoadException.hpp" />
    <ClInclude Include="gel\asset\FontEncode.hpp" />
    <ClInclude Include="gel\asset\FontLayout.hpp" />
    <ClInclude Include="gel\asset\FreeTypeFont.hpp" />
    <ClInclude Include="gel\asset\FreeTypeFontInstance.hpp" />
    <ClInclude Include="gel\asset\IFont.hpp" />
    <ClInclude Include="gel\asset\IFontInstance.hpp" />
    <ClInclude Include="gel\asset\ProxyTexture.hpp" />
    <ClInclude Include="gel\GameConfig.hpp" />
    <ClInclude Include="gel\input\ButtonState.hpp" />
    <ClInclude Include="gel\input\InputState.hpp" />
    <ClInclude Include="gel\input\InputSystem.hpp" />
    <ClInclude Include="gel\input\KeyboardBuffer.hpp" />
    <ClInclude Include="gel\input\KeyboardState.hpp" />
    <ClInclude Include="gel\input\MouseBuffer.hpp" />
    <ClInclude Include="gel\input\MouseButton.hpp" />
    <ClInclude Include="gel\input\MouseState.hpp" />
    <ClInclude Include="gel\pipeline\TtfPipeline.hpp" />
    <ClInclude Include="gel\freetype.hpp" />
    <ClInclude Include="gel\input\Drag.hpp" />
    <ClInclude Include="gel\input\Input.hpp" />
    <ClInclude Include="gel\pipeline\BmpPipeline.hpp" />
    <ClInclude Include="gel\pipeline\ContentManager.hpp" />
    <ClInclude Include="gel\pipeline\FbxPipeline.hpp" />
    <ClInclude Include="gel\pipeline\IContentPipeline.hpp" />
    <ClInclude Include="gel\pipeline\JpegPipeline.hpp" />
    <ClInclude Include="gel\pipeline\OtfPipeline.hpp" />
    <ClInclude Include="gel\pipeline\PngPipeline.hpp" />
    <ClInclude Include="gel\pipeline\ProxyPipeline.hpp" />
    <ClInclude Include="gel\pipeline\WavePipeline.hpp" />
    <ClInclude Include="gel\asset\AssetDatabase.hpp" />
    <ClInclude Include="gel\asset\FbxModel.hpp" />
    <ClInclude Include="gel\asset\fbxsdk.h" />
    <ClInclude Include="gel\asset\IAsset.hpp" />
    <ClInclude Include="gel\asset\IModel.hpp" />
    <ClInclude Include="gel\asset\IAudio.hpp" />
    <ClInclude Include="gel\asset\ITexture.hpp" />
    <ClInclude Include="gel\asset\SOILTexture.hpp" />
    <ClInclude Include="gel\asset\TextureIO.hpp" />
    <ClInclude Include="gel\asset\WaveAudio.hpp" />
    <ClInclude Include="gel\fbxsdk.hpp" />
    <ClInclude Include="gel\Game.hpp" />
    <ClInclude Include="gel\gel.hpp" />
    <ClInclude Include="gel\gli.hpp" />
    <ClInclude Include="gel\gsystem\Camera.hpp" />
    <ClInclude Include="gel\gsystem\Duration.hpp" />
    <ClInclude Include="gel\gsystem\LineRenderer.hpp" />
    <ClInclude Include="gel\math\Random.hpp" />
    <ClInclude Include="gel\gsystem\Timer.hpp" />
    <ClInclude Include="gel\gsystem\Transform.hpp" />
    <ClInclude Include="gel\math\AABB.hpp" />
    <ClInclude Include="gel\math\Quadrangle.hpp" />
    <ClInclude Include="gel\math\Rectangle.hpp" />
    <ClInclude Include="gel\math\Triangle.hpp" />
    <ClInclude Include="gel\scene\AnimationComponent.hpp" />
    <ClInclude Include="gel\scene\Component.hpp" />
    <ClInclude Include="gel\scene\Disposable.hpp" />
    <ClInclude Include="gel\scene\GameObject.hpp" />
    <ClInclude Include="gel\scene\RendererComponent.hpp" />
    <ClInclude Include="gel\scene\Scene.hpp" />
    <ClInclude Include="gel\scene\SceneManager.hpp" />
    <ClInclude Include="gel\scene\SpriteComponent.hpp" />
    <ClInclude Include="gel\scene\TimerComponent.hpp" />
    <ClInclude Include="gel\scene\UIArrowButtonComponent.hpp" />
    <ClInclude Include="gel\scene\UIButtonComponent.hpp" />
    <ClInclude Include="gel\scene\UIElementComponent.hpp" />
    <ClInclude Include="gel\scene\UIFrameComponent.hpp" />
    <ClInclude Include="gel\scene\UIImageComponent.hpp" />
    <ClInclude Include="gel\scene\UILayoutComponent.hpp" />
    <ClInclude Include="gel\scene\UIProgressBarComponent.hpp" />
    <ClInclude Include="gel\service\FBXSDKService.hpp" />
    <ClInclude Include="gel\service\FreeTypeService.hpp" />
    <ClInclude Include="gel\service\IService.hpp" />
    <ClInclude Include="gel\shader\Renderer.hpp" />
    <ClInclude Include="gel\signal\Property.hpp" />
    <ClInclude Include="gel\scene\UIRelativeLayoutComponent.hpp" />
    <ClInclude Include="gel\scene\UITextComponent.hpp" />
    <ClInclude Include="gel\scene\UITheme.hpp" />
    <ClInclude Include="gel\shader\Bitmap.hpp" />
    <ClInclude Include="gel\shader\Box.hpp" />
    <ClInclude Include="gel\shader\Buffer.hpp" />
    <ClInclude Include="gel\shader\Circle.hpp" />
    <ClInclude Include="gel\shader\Color4.hpp" />
    <ClInclude Include="gel\shader\CubeMap.hpp" />
    <ClInclude Include="gel\shader\FontRenderer.hpp" />
    <ClInclude Include="gel\shader\FrameBuffer.hpp" />
    <ClInclude Include="gel\shader\IRMaterial.hpp" />
    <ClInclude Include="gel\shader\IRMesh.hpp" />
    <ClInclude Include="gel\shader\IRModel.hpp" />
    <ClInclude Include="gel\shader\IRShape.hpp" />
    <ClInclude Include="gel\shader\Layer.hpp" />
    <ClInclude Include="gel\shader\Line.hpp" />
    <ClInclude Include="gel\shader\NameRule.hpp" />
    <ClInclude Include="gel\shader\PixelBuffer.hpp" />
    <ClInclude Include="gel\shader\Plane.hpp" />
    <ClInclude Include="gel\shader\PlaneBatch.hpp" />
    <ClInclude Include="gel\shader\PlaneDraw.hpp" />
    <ClInclude Include="gel\shader\Primitive.hpp" />
    <ClInclude Include="gel\shader\RawTexture.hpp" />
    <ClInclude Include="gel\shader\Reflection4.hpp" />
    <ClInclude Include="gel\shader\RenderBuffer.hpp" />
    <ClInclude Include="gel\shader\ScreenBuffer.hpp" />
    <ClInclude Include="gel\shader\Sequence.hpp" />
    <ClInclude Include="gel\shader\Shader.hpp" />
    <ClInclude Include="gel\shader\ShaderRegistry.hpp" />
    <ClInclude Include="gel\shader\VertexArray.hpp" />
    <ClInclude Include="gel\signal\Connection.hpp" />
    <ClInclude Include="gel\signal\IConnectable.hpp" />
    <ClInclude Include="gel\signal\Signal.hpp" />
    <ClInclude Include="gel\signal\Slot.hpp" />
    <ClInclude Include="gel\pipeline\Thread.hpp" />
    <ClInclude Include="gel\util\BlockingQueue.hpp" />
    <ClInclude Include="gel\util\collections.hpp" />
    <ClInclude Include="gel\util\CommandReader.hpp" />
    <ClInclude Include="gel\util\enum.hpp" />
    <ClInclude Include="gel\util\error.hpp" />
    <ClInclude Include="gel\util\Flag.hpp" />
    <ClInclude Include="gel\util\io.hpp" />
    <ClInclude Include="gel\util\KeyTrigger.hpp" />
    <ClInclude Include="gel\util\MouseTrigger.hpp" />
    <ClInclude Include="gel\util\Names.hpp" />
    <ClInclude Include="gel\util\NonCopyable.hpp" />
    <ClInclude Include="gel\util\SharedHelper.hpp" />
    <ClInclude Include="gel\util\string.hpp" />
    <ClInclude Include="gel\util\Vec2HashFunc.hpp" />
    <ClInclude Include="gel\Window.hpp" />
  </ItemGroup>
  <ItemGroup>
    <None Include="gel\scene\UIRect.hpp" />
    <None Include="packages.config" />
  </ItemGroup>
  <Import Project="$(VCTargetsPath)\Microsoft.Cpp.targets" />
  <ImportGroup Label="ExtensionTargets">
    <Import Project="..\packages\freeglut.redist.2.8.1.15\build\native\freeglut.redist.targets" Condition="Exists('..\packages\freeglut.redist.2.8.1.15\build\native\freeglut.redist.targets')" />
    <Import Project="..\packages\freeglut.2.8.1.15\build\native\freeglut.targets" Condition="Exists('..\packages\freeglut.2.8.1.15\build\native\freeglut.targets')" />
    <Import Project="..\packages\glew.redist.1.9.0.1\build\native\glew.redist.targets" Condition="Exists('..\packages\glew.redist.1.9.0.1\build\native\glew.redist.targets')" />
    <Import Project="..\packages\glew.1.9.0.1\build\native\glew.targets" Condition="Exists('..\packages\glew.1.9.0.1\build\native\glew.targets')" />
    <Import Project="..\packages\glfw.3.3.0.1\build\native\glfw.targets" Condition="Exists('..\packages\glfw.3.3.0.1\build\native\glfw.targets')" />
    <Import Project="..\packages\glm.0.9.9.600\build\native\glm.targets" Condition="Exists('..\packages\glm.0.9.9.600\build\native\glm.targets')" />
    <Import Project="..\packages\freetype.redist.2.8.0.1\build\native\freetype.redist.targets" Condition="Exists('..\packages\freetype.redist.2.8.0.1\build\native\freetype.redist.targets')" />
    <Import Project="..\packages\freetype.2.8.0.1\build\native\freetype.targets" Condition="Exists('..\packages\freetype.2.8.0.1\build\native\freetype.targets')" />
  </ImportGroup>
  <Target Name="EnsureNuGetPackageBuildImports" BeforeTargets="PrepareForBuild">
    <PropertyGroup>
      <ErrorText>このプロジェクトは、このコンピューター上にない NuGet パッケージを参照しています。それらのパッケージをダウンロードするには、[NuGet パッケージの復元] を使用します。詳細については、http://go.microsoft.com/fwlink/?LinkID=322105 を参照してください。見つからないファイルは {0} です。</ErrorText>
    </PropertyGroup>
    <Error Condition="!Exists('..\packages\freeglut.redist.2.8.1.15\build\native\freeglut.redist.targets')" Text="$([System.String]::Format('$(ErrorText)', '..\packages\freeglut.redist.2.8.1.15\build\native\freeglut.redist.targets'))" />
    <Error Condition="!Exists('..\packages\freeglut.2.8.1.15\build\native\freeglut.targets')" Text="$([System.String]::Format('$(ErrorText)', '..\packages\freeglut.2.8.1.15\build\native\freeglut.targets'))" />
    <Error Condition="!Exists('..\packages\glew.redist.1.9.0.1\build\native\glew.redist.targets')" Text="$([System.String]::Format('$(ErrorText)', '..\packages\glew.redist.1.9.0.1\build\native\glew.redist.targets'))" />
    <Error Condition="!Exists('..\packages\glew.1.9.0.1\build\native\glew.targets')" Text="$([System.String]::Format('$(ErrorText)', '..\packages\glew.1.9.0.1\build\native\glew.targets'))" />
    <Error Condition="!Exists('..\packages\glfw.3.3.0.1\build\native\glfw.targets')" Text="$([System.String]::Format('$(ErrorText)', '..\packages\glfw.3.3.0.1\build\native\glfw.targets'))" />
    <Error Condition="!Exists('..\packages\glm.0.9.9.600\build\native\glm.targets')" Text="$([System.String]::Format('$(ErrorText)', '..\packages\glm.0.9.9.600\build\native\glm.targets'))" />
    <Error Condition="!Exists('..\packages\soil.1.16.0\build\native\soil.props')" Text="$([System.String]::Format('$(ErrorText)', '..\packages\soil.1.16.0\build\native\soil.props'))" />
    <Error Condition="!Exists('..\packages\freetype.redist.2.8.0.1\build\native\freetype.redist.targets')" Text="$([System.String]::Format('$(ErrorText)', '..\packages\freetype.redist.2.8.0.1\build\native\freetype.redist.targets'))" />
    <Error Condition="!Exists('..\packages\freetype.2.8.0.1\build\native\freetype.targets')" Text="$([System.String]::Format('$(ErrorText)', '..\packages\freetype.2.8.0.1\build\native\freetype.targets'))" />
  </Target>
</Project>
````
**OUTPUT**
````
package main
type ProjectItemDefinitionGroup struct {
    // define attribute
    AttrCondition string `xml:"Condition,attr"`
    // define subelement
    SubClCompile []*ProjectItemDefinitionGroupClCompile `xml:"ClCompile"`
    SubLink []*ProjectItemDefinitionGroupLink `xml:"Link"`
    SubPreBuildEvent []*ProjectItemDefinitionGroupPreBuildEvent `xml:"PreBuildEvent"`
    // define content
    Content string `xml:",chardata"`
}
type ProjectPropertyGroupPlatformToolset struct {
    // define attribute
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectItemDefinitionGroupPreBuildEvent struct {
    // define attribute
    // define subelement
    SubCommand []*ProjectItemDefinitionGroupPreBuildEventCommand `xml:"Command"`
    // define content
    Content string `xml:",chardata"`
}
type ProjectItemDefinitionGroupLinkAdditionalLibraryDirectories struct {
    // define attribute
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectTargetPropertyGroup struct {
    // define attribute
    // define subelement
    SubErrorText []*ProjectTargetPropertyGroupErrorText `xml:"ErrorText"`
    // define content
    Content string `xml:",chardata"`
}
type ProjectItemGroupClInclude struct {
    // define attribute
    AttrInclude string `xml:"Include,attr"`
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectImport struct {
    // define attribute
    AttrCondition string `xml:"Condition,attr"`
    AttrProject string `xml:"Project,attr"`
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectItemGroup struct {
    // define attribute
    AttrLabel string `xml:"Label,attr"`
    // define subelement
    SubProjectConfiguration []*ProjectItemGroupProjectConfiguration `xml:"ProjectConfiguration"`
    SubClCompile []*ProjectItemGroupClCompile `xml:"ClCompile"`
    SubClInclude []*ProjectItemGroupClInclude `xml:"ClInclude"`
    SubNone []*ProjectItemGroupNone `xml:"None"`
    // define content
    Content string `xml:",chardata"`
}
type ProjectPropertyGroupUseDebugLibraries struct {
    // define attribute
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectPropertyGroupWholeProgramOptimization struct {
    // define attribute
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectItemDefinitionGroupClCompileRuntimeLibrary struct {
    // define attribute
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectItemDefinitionGroupLinkAdditionalDependencies struct {
    // define attribute
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectPropertyGroupRootNamespace struct {
    // define attribute
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectImportGroupImport struct {
    // define attribute
    AttrProject string `xml:"Project,attr"`
    AttrCondition string `xml:"Condition,attr"`
    AttrLabel string `xml:"Label,attr"`
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectItemDefinitionGroupLink struct {
    // define attribute
    // define subelement
    SubAdditionalLibraryDirectories []*ProjectItemDefinitionGroupLinkAdditionalLibraryDirectories `xml:"AdditionalLibraryDirectories"`
    SubAdditionalDependencies []*ProjectItemDefinitionGroupLinkAdditionalDependencies `xml:"AdditionalDependencies"`
    SubEnableCOMDATFolding []*ProjectItemDefinitionGroupLinkEnableCOMDATFolding `xml:"EnableCOMDATFolding"`
    SubOptimizeReferences []*ProjectItemDefinitionGroupLinkOptimizeReferences `xml:"OptimizeReferences"`
    // define content
    Content string `xml:",chardata"`
}
type ProjectItemGroupNone struct {
    // define attribute
    AttrInclude string `xml:"Include,attr"`
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectTargetError struct {
    // define attribute
    AttrCondition string `xml:"Condition,attr"`
    AttrText string `xml:"Text,attr"`
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectPropertyGroup struct {
    // define attribute
    AttrLabel string `xml:"Label,attr"`
    AttrCondition string `xml:"Condition,attr"`
    // define subelement
    SubVCProjectVersion []*ProjectPropertyGroupVCProjectVersion `xml:"VCProjectVersion"`
    SubProjectGuid []*ProjectPropertyGroupProjectGuid `xml:"ProjectGuid"`
    SubRootNamespace []*ProjectPropertyGroupRootNamespace `xml:"RootNamespace"`
    SubWindowsTargetPlatformVersion []*ProjectPropertyGroupWindowsTargetPlatformVersion `xml:"WindowsTargetPlatformVersion"`
    SubConfigurationType []*ProjectPropertyGroupConfigurationType `xml:"ConfigurationType"`
    SubUseDebugLibraries []*ProjectPropertyGroupUseDebugLibraries `xml:"UseDebugLibraries"`
    SubPlatformToolset []*ProjectPropertyGroupPlatformToolset `xml:"PlatformToolset"`
    SubCharacterSet []*ProjectPropertyGroupCharacterSet `xml:"CharacterSet"`
    SubWholeProgramOptimization []*ProjectPropertyGroupWholeProgramOptimization `xml:"WholeProgramOptimization"`
    // define content
    Content string `xml:",chardata"`
}
type ProjectItemGroupProjectConfiguration struct {
    // define attribute
    AttrInclude string `xml:"Include,attr"`
    // define subelement
    SubConfiguration []*ProjectItemGroupProjectConfigurationConfiguration `xml:"Configuration"`
    SubPlatform []*ProjectItemGroupProjectConfigurationPlatform `xml:"Platform"`
    // define content
    Content string `xml:",chardata"`
}
type ProjectPropertyGroupCharacterSet struct {
    // define attribute
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectItemDefinitionGroupClCompileOptimization struct {
    // define attribute
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectItemDefinitionGroupClCompilePreprocessorDefinitions struct {
    // define attribute
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type Project struct {
    // define attribute
    Attrxmlns string `xml:"xmlns,attr"`
    AttrDefaultTargets string `xml:"DefaultTargets,attr"`
    AttrToolsVersion string `xml:"ToolsVersion,attr"`
    // define subelement
    SubImport []*ProjectImport `xml:"Import"`
    SubItemGroup []*ProjectItemGroup `xml:"ItemGroup"`
    SubPropertyGroup []*ProjectPropertyGroup `xml:"PropertyGroup"`
    SubImportGroup []*ProjectImportGroup `xml:"ImportGroup"`
    SubItemDefinitionGroup []*ProjectItemDefinitionGroup `xml:"ItemDefinitionGroup"`
    SubTarget []*ProjectTarget `xml:"Target"`
    // define content
    Content string `xml:",chardata"`
}
type ProjectTarget struct {
    // define attribute
    AttrBeforeTargets string `xml:"BeforeTargets,attr"`
    AttrName string `xml:"Name,attr"`
    // define subelement
    SubPropertyGroup []*ProjectTargetPropertyGroup `xml:"PropertyGroup"`
    SubError []*ProjectTargetError `xml:"Error"`
    // define content
    Content string `xml:",chardata"`
}
type ProjectItemGroupProjectConfigurationPlatform struct {
    // define attribute
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectPropertyGroupProjectGuid struct {
    // define attribute
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectItemDefinitionGroupClCompile struct {
    // define attribute
    // define subelement
    SubWarningLevel []*ProjectItemDefinitionGroupClCompileWarningLevel `xml:"WarningLevel"`
    SubOptimization []*ProjectItemDefinitionGroupClCompileOptimization `xml:"Optimization"`
    SubSDLCheck []*ProjectItemDefinitionGroupClCompileSDLCheck `xml:"SDLCheck"`
    SubConformanceMode []*ProjectItemDefinitionGroupClCompileConformanceMode `xml:"ConformanceMode"`
    SubAdditionalIncludeDirectories []*ProjectItemDefinitionGroupClCompileAdditionalIncludeDirectories `xml:"AdditionalIncludeDirectories"`
    SubLanguageStandard []*ProjectItemDefinitionGroupClCompileLanguageStandard `xml:"LanguageStandard"`
    SubPreprocessorDefinitions []*ProjectItemDefinitionGroupClCompilePreprocessorDefinitions `xml:"PreprocessorDefinitions"`
    SubRuntimeLibrary []*ProjectItemDefinitionGroupClCompileRuntimeLibrary `xml:"RuntimeLibrary"`
    SubFunctionLevelLinking []*ProjectItemDefinitionGroupClCompileFunctionLevelLinking `xml:"FunctionLevelLinking"`
    SubIntrinsicFunctions []*ProjectItemDefinitionGroupClCompileIntrinsicFunctions `xml:"IntrinsicFunctions"`
    // define content
    Content string `xml:",chardata"`
}
type ProjectItemDefinitionGroupLinkOptimizeReferences struct {
    // define attribute
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectPropertyGroupVCProjectVersion struct {
    // define attribute
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectPropertyGroupWindowsTargetPlatformVersion struct {
    // define attribute
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectItemDefinitionGroupClCompileWarningLevel struct {
    // define attribute
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectItemDefinitionGroupClCompileLanguageStandard struct {
    // define attribute
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectItemDefinitionGroupPreBuildEventCommand struct {
    // define attribute
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectTargetPropertyGroupErrorText struct {
    // define attribute
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectImportGroup struct {
    // define attribute
    AttrCondition string `xml:"Condition,attr"`
    AttrLabel string `xml:"Label,attr"`
    // define subelement
    SubImport []*ProjectImportGroupImport `xml:"Import"`
    // define content
    Content string `xml:",chardata"`
}
type ProjectPropertyGroupConfigurationType struct {
    // define attribute
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectItemDefinitionGroupClCompileConformanceMode struct {
    // define attribute
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectItemDefinitionGroupClCompileAdditionalIncludeDirectories struct {
    // define attribute
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectItemDefinitionGroupClCompileFunctionLevelLinking struct {
    // define attribute
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectItemDefinitionGroupLinkEnableCOMDATFolding struct {
    // define attribute
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectItemGroupProjectConfigurationConfiguration struct {
    // define attribute
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectItemDefinitionGroupClCompileSDLCheck struct {
    // define attribute
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectItemDefinitionGroupClCompileIntrinsicFunctions struct {
    // define attribute
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
type ProjectItemGroupClCompile struct {
    // define attribute
    AttrInclude string `xml:"Include,attr"`
    // define subelement
    // define content
    Content string `xml:",chardata"`
}
func LoadProject(path string) (*Project, error) {
    xmlFile, err := os.Open(path)
    if err != nil {
    	return nil, err
    }
    defer xmlFile.Close()
    xmlData, err := ioutil.ReadAll(xmlFile)
    if err != nil {
    	return nil, err
    }
    var data Project
    xml.Unmarshal(xmlData, &data)
    return &data, nil
}

func SaveProject(path string, data *Project, perm os.FileMode) error {
    buf, err := xml.MarshalIndent(data, "", "    ")
    if err != nil {
        return err
    }
    err = ioutil.WriteFile(path, buf, perm)
    if err != nil {
    	return err
    }
    return nil
}

````