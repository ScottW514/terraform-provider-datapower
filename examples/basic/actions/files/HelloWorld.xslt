<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">
    <xsl:output method="html" encoding="UTF-8" doctype-public="-//W3C//DTD HTML 4.01//EN" doctype-system="http://www.w3.org/TR/html4/strict.dtd" indent="yes"/>

    <xsl:template match="/">
        <html>
            <head>
                <title>Hello World Page</title>
            </head>
            <body>
                <h1><xsl:value-of select="root/message"/></h1>
                <p>This is a simple Hello World web page generated from XML via XSLT transformation.</p>
            </body>
        </html>
    </xsl:template>
</xsl:stylesheet>
